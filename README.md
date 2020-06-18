# Mercantile:  An EarthCube Graph Server

## Status
This is very early work!  Some issues include:

* The options heavy SPARQL is SLOOOOW   need to code in logic to limit the options 
* The query offset is not wired in...  just need to do that. 
* The query pattern is not resolved...  need to work up how to map interfaces query 
needs to te GraphQL pattern. 

## About
This is a GraphQL server that connects into a defined RDF triple store.  
It fronts a set of SPARQL calls to a data interface defined from the
schema.org type Dataset "class" (ref: https://schema.org/Dataset).  

This is a simple code gen build based on the 99 Designs gqlgen package
(ref: https://github.com/99designs/gqlgen).    This leverages the 
schema.org types as interface guidance.  My desire is to use SHACL shape
graphs for this.  This has been done by TopQuadrant but is obviously 
closed source.  Shape to interface should be easy if I get time.  

## Mercantile ?
I'm not fully convinced that GraphQL is the best way and it's obviously 
not the only way.  Other options 
being JSON:API for example.  Also this obviously doesn't address other 
access methods like KoopJS or S3 Query patterns.  Mercantile seems a 
sufficiently generic name to scope this work under.  
