package dbase

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/earthcubearchitecture-project418/mercantile/graph/model"
	"github.com/knakk/sparql"
	"github.com/spf13/viper"
)

const queries = `
# Comments are ignored, except those tagging a query.

# tag: generalInfo
PREFIX text: <http://jena.apache.org/text#>
PREFIX  schema: <http://schema.org/>

SELECT ?name ?keywords ?distribution ?description
WHERE {
   { GRAPH ?g {
	   ?s schema:description ?description .
	   ?s schema:name ?name .
	   ?s schema:keywords ?keywords .
	   ?s schema:url ?distribution .
	   }
	}
}
LIMIT 10

# removed      OPTIONAL {?s schema:additionalType ?addtype . }   from the following
# may need SELECT DISTINCT * if we use template controlled optional elements
#tag: mainsearch
prefix schema: <https://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
SELECT DISTINCT ?subj  ?p ?score  ?type  ?name ?relto ?addtype ?url  ?description
 WHERE {
   ?lit bds:search "{{.Q}}" .
   {{ if .MatchAll}}
   ?lit bds:matchAllTerms "false" .
   {{end}}
   ?lit bds:relevance ?score .
   ?subj ?p ?lit .
   BIND (?subj as ?s)
   OPTIONAL {?s rdf:type ?type .}
   OPTIONAL {?s schema:additionalType ?addtype . }
   OPTIONAL {?s schema:name ?name .}
   OPTIONAL {?s schema:isRelatedTo ?relto .}
   OPTIONAL {?s schema:url ?url . }
   OPTIONAL {?s schema:description ?description . }
 }
ORDER BY DESC(?score)
OFFSET 0


# removed      OPTIONAL {?s schema:additionalType ?addtype . }
# and relto from the following
# may need SELECT DISTINCT * if we use template controlled optional elements
# NOTE hack that locks this to type dataset (then stupidly does an OPTIONAL for that  )  ;)
#tag: geodex
prefix schema: <http://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
SELECT DISTINCT ?subj ?p ?score  ?type  ?name ?addtype ?url  ?description
 WHERE {
   ?lit bds:search "{{.Q}}" .
   {{ if .MatchAll}}
   ?lit bds:matchAllTerms "false" .
   {{end}}
   ?lit bds:relevance ?score .
   ?subj ?p ?lit .
   BIND (?subj as ?s)
   ?s a schema:Dataset .
   OPTIONAL {?s schema:distribution ?dis .
	?dis schema:url ?url . }
   OPTIONAL {?s rdf:type ?type .}
   OPTIONAL {?s schema:additionalType ?addtype . }
   OPTIONAL {?s schema:name ?name .}
   OPTIONAL {?s schema:description ?description . }
 }
ORDER BY DESC(?score) ?subj
OFFSET {{.Offset}}
LIMIT {{.First}}

#tag: group
prefix schema: <http://schema.org/>
prefix bds: <http://www.bigdata.com/rdf/search#>
SELECT DISTINCT ?subj ?url ?score ?name ?description
 WHERE {
   ?lit bds:search "{{.Q}}" .
   {{ if .MatchAll}}
   ?lit bds:matchAllTerms "false" .
   {{end}}
   ?lit bds:relevance ?score .
   ?subj ?p ?lit .
   ?subj a schema:Dataset .
   BIND (?subj as ?s)
   {
   		SELECT  ?s (MIN(?disurl) as ?url) {
            ?s a schema:Dataset .
            ?s schema:distribution ?dis .
   			?dis schema:url ?disurl .
   	  	} GROUP BY ?s
   }
   ?s schema:name ?name .
   ?s schema:description ?description .
 }
ORDER BY DESC(?score) ?subj
OFFSET {{.Offset}}
LIMIT {{.First}}

`

// QModel holds the query and all word matching request boolean
type QModel struct {
	Q        string // query string
	MatchAll bool   // do we match all terms
	First    int    // number to get
	Offset   int    // where to start "geting" from in the list
}

var viperVal string

// DescriptionCall make a SPARQL call
func DescriptionCall(q, u string, first, offset int) ([]*model.Do, error) {
	// Viper config   # should be passed as a pointer
	var v1 *viper.Viper
	v1, err := readConfig(viperVal, nil)
	if err != nil {
		panic(fmt.Errorf("error when reading config: %v", err))
	}

	mcfg := v1.GetStringMapString("server")
	log.Println(mcfg["endpoint"])
	repo, err := getLocalSPARQL(mcfg["endpoint"])

	// At this point we need to route the request to the queries that support the request.  This will be based on the request parameters sent
	log.Printf("q: %s", q)
	log.Printf("u: %s", u)

	f := bytes.NewBufferString(queries)
	bank := sparql.LoadBank(f)

	qm := QModel{Q: q, MatchAll: true, First: first, Offset: offset}

	sq, err := bank.Prepare("group", qm) // change the SPARQL call used here..   mainsearch is the other one for now
	if err != nil {
		log.Printf("%s\n", err)
	}

	log.Println(sq)

	res, err := repo.Query(sq)
	if err != nil {
		log.Printf("query call: %v\n", err)
		return nil, err
	}

	// set up our struct
	ds := []*model.Do{}

	bindingsTest1 := res.Solutions() // []map[string]red.Term

	for x := range bindingsTest1 {
		d := &model.Do{}
		if val, ok := bindingsTest1[x]["subj"]; ok {
			d.Subject = val.String()
		}
		if val, ok := bindingsTest1[x]["score"]; ok {
			d.Score = val.String()
		}
		if val, ok := bindingsTest1[x]["type"]; ok {
			d.Type = val.String()
		}
		if val, ok := bindingsTest1[x]["name"]; ok {
			d.Name = val.String()
		}
		if val, ok := bindingsTest1[x]["relto"]; ok {
			d.Relto = val.String()
		}
		if val, ok := bindingsTest1[x]["addtype"]; ok {
			d.Addtype = val.String()
		}
		if val, ok := bindingsTest1[x]["url"]; ok {
			d.URL = val.String()
		}
		if val, ok := bindingsTest1[x]["description"]; ok {
			d.Description = val.String()
		}
		ds = append(ds, d)
	}

	/*
		bindingsTest2 := res.Bindings() // map[string][]rdf.Term
		log.Printf("Binding Test %s ", bindingsTest2)
		log.Println(len(bindingsTest2["description"]))

		// This whole aspect seems verbose... there has to be a better Go way to do this check?
		if len(bindingsTest2) > 0 {
			if len(bindingsTest2["description"]) > 0 {
				for x := range bindingsTest2["description"] {
					d := &model.Do{}
					d.Name = bindingsTest2["name"][x].String()
					d.Description = bindingsTest2["description"][x].String()
					d.URL = bindingsTest2["url"][x].String()
					d.Score = bindingsTest2["score"][x].String()
					ds = append(ds, d)
				}
			}
		}
	*/

	return ds, err
}

// connector function for the local sparql instance
func getLocalSPARQL(endpoint string) (*sparql.Repo, error) {
	repo, err := sparql.NewRepo(endpoint,
		sparql.Timeout(time.Millisecond*30000),
	)
	if err != nil {
		log.Printf("getLocalSPARQL %s\n", err)
	}
	return repo, err
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
