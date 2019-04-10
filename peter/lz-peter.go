package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"github.com/andrew-d/lzma"
	"io/ioutil"
	"math"
	"time"
)

func main() {
	//	data := `{
	//  "sessionAttributes": {
	//    "afm": "fare_amount_sum",
	//    "afmBusinessNames": "{\"fare_amount\":\"Total Fare Amount\"}",
	//    "bitmap": "110101011110",
	//    "graph": "{\"title\":\"for Passenger Count 0 for Pickup Location Zone Allerton/Pelham Gardens by Drop-off Location Borough over Pickup Date\",\"series\":[\"Bronx\",\"Brooklyn\",\"EWR\",\"Manhattan\",\"Queens\",\"Staten Island\",\"Unknown\"],\"vids\":[\"0\"],\"seriesVids\":{\"Bronx\":\"0\"},\"atms\":[\"0\"],\"templateValues\":[\"0\",\"Allerton/Pelham Gardens\"]}",
	//    "graphTemplate": "for Passenger Count %s for Pickup Location Zone %s by Drop-off Location Borough over Pickup Date",
	//    "graphTitleValues": "[[\"0\",\"Allerton/Pelham Gardens\"],[\"1\",\"Allerton/Pelham Gardens\"],[\"2\",\"Allerton/Pelham Gardens\"],[\"3\",\"Allerton/Pelham Gardens\"],[\"4\",\"Allerton/Pelham Gardens\"],[\"5\",\"Allerton/Pelham Gardens\"],[\"6\",\"Allerton/Pelham Gardens\"],[\"7\",\"Allerton/Pelham Gardens\"],[\"8\",\"Allerton/Pelham Gardens\"],[\"9\",\"Allerton/Pelham Gardens\"],[\"0\",\"Alphabet City\"],[\"1\",\"Alphabet City\"],[\"2\",\"Alphabet City\"],[\"3\",\"Alphabet City\"],[\"4\",\"Alphabet City\"],[\"5\",\"Alphabet City\"],[\"6\",\"Alphabet City\"],[\"7\",\"Alphabet City\"],[\"8\",\"Alphabet City\"],[\"9\",\"Alphabet City\"],[\"0\",\"Arden Heights\"],[\"1\",\"Arden Heights\"],[\"2\",\"Arden Heights\"],[\"3\",\"Arden Heights\"],[\"4\",\"Arden Heights\"],[\"5\",\"Arden Heights\"],[\"6\",\"Arden Heights\"],[\"7\",\"Arden Heights\"],[\"8\",\"Arden Heights\"],[\"9\",\"Arden Heights\"],[\"0\",\"Arrochar/Fort Wadsworth\"],[\"1\",\"Arrochar/Fort Wadsworth\"],[\"2\",\"Arrochar/Fort Wadsworth\"],[\"3\",\"Arrochar/Fort Wadsworth\"],[\"4\",\"Arrochar/Fort Wadsworth\"],[\"5\",\"Arrochar/Fort Wadsworth\"],[\"6\",\"Arrochar/Fort Wadsworth\"],[\"7\",\"Arrochar/Fort Wadsworth\"],[\"8\",\"Arrochar/Fort Wadsworth\"],[\"9\",\"Arrochar/Fort Wadsworth\"],[\"0\",\"Astoria\"],[\"1\",\"Astoria\"],[\"2\",\"Astoria\"],[\"3\",\"Astoria\"],[\"4\",\"Astoria\"],[\"5\",\"Astoria\"],[\"6\",\"Astoria\"],[\"7\",\"Astoria\"],[\"8\",\"Astoria\"],[\"9\",\"Astoria\"],[\"0\",\"Astoria Park\"],[\"1\",\"Astoria Park\"],[\"2\",\"Astoria Park\"],[\"3\",\"Astoria Park\"],[\"4\",\"Astoria Park\"],[\"5\",\"Astoria Park\"],[\"6\",\"Astoria Park\"],[\"7\",\"Astoria Park\"],[\"8\",\"Astoria Park\"],[\"9\",\"Astoria Park\"],[\"0\",\"Auburndale\"],[\"1\",\"Auburndale\"],[\"2\",\"Auburndale\"],[\"3\",\"Auburndale\"],[\"4\",\"Auburndale\"],[\"5\",\"Auburndale\"],[\"6\",\"Auburndale\"],[\"7\",\"Auburndale\"],[\"8\",\"Auburndale\"],[\"9\",\"Auburndale\"],[\"0\",\"Baisley Park\"],[\"1\",\"Baisley Park\"],[\"2\",\"Baisley Park\"],[\"3\",\"Baisley Park\"],[\"4\",\"Baisley Park\"],[\"5\",\"Baisley Park\"],[\"6\",\"Baisley Park\"],[\"7\",\"Baisley Park\"],[\"8\",\"Baisley Park\"],[\"9\",\"Baisley Park\"],[\"0\",\"Bath Beach\"],[\"1\",\"Bath Beach\"],[\"2\",\"Bath Beach\"],[\"3\",\"Bath Beach\"],[\"4\",\"Bath Beach\"],[\"5\",\"Bath Beach\"],[\"6\",\"Bath Beach\"],[\"7\",\"Bath Beach\"],[\"8\",\"Bath Beach\"],[\"9\",\"Bath Beach\"],[\"0\",\"Battery Park\"],[\"1\",\"Battery Park\"],[\"2\",\"Battery Park\"],[\"3\",\"Battery Park\"],[\"4\",\"Battery Park\"],[\"5\",\"Battery Park\"],[\"6\",\"Battery Park\"],[\"7\",\"Battery Park\"],[\"8\",\"Battery Park\"],[\"9\",\"Battery Park\"],[\"0\",\"Battery Park City\"],[\"1\",\"Battery Park City\"],[\"2\",\"Battery Park City\"],[\"3\",\"Battery Park City\"],[\"4\",\"Battery Park City\"],[\"5\",\"Battery Park City\"],[\"6\",\"Battery Park City\"],[\"7\",\"Battery Park City\"],[\"8\",\"Battery Park City\"],[\"9\",\"Battery Park City\"],[\"0\",\"Bay Ridge\"],[\"1\",\"Bay Ridge\"],[\"2\",\"Bay Ridge\"],[\"3\",\"Bay Ridge\"],[\"4\",\"Bay Ridge\"],[\"5\",\"Bay Ridge\"],[\"6\",\"Bay Ridge\"],[\"7\",\"Bay Ridge\"],[\"8\",\"Bay Ridge\"],[\"9\",\"Bay Ridge\"],[\"0\",\"Bay Terrace/Fort Totten\"],[\"1\",\"Bay Terrace/Fort Totten\"],[\"2\",\"Bay Terrace/Fort Totten\"],[\"3\",\"Bay Terrace/Fort Totten\"],[\"4\",\"Bay Terrace/Fort Totten\"],[\"5\",\"Bay Terrace/Fort Totten\"],[\"6\",\"Bay Terrace/Fort Totten\"],[\"7\",\"Bay Terrace/Fort Totten\"],[\"8\",\"Bay Terrace/Fort Totten\"],[\"9\",\"Bay Terrace/Fort Totten\"],[\"0\",\"Bayside\"],[\"1\",\"Bayside\"],[\"2\",\"Bayside\"],[\"3\",\"Bayside\"],[\"4\",\"Bayside\"],[\"5\",\"Bayside\"],[\"6\",\"Bayside\"],[\"7\",\"Bayside\"],[\"8\",\"Bayside\"],[\"9\",\"Bayside\"],[\"0\",\"Bedford\"],[\"1\",\"Bedford\"],[\"2\",\"Bedford\"],[\"3\",\"Bedford\"],[\"4\",\"Bedford\"],[\"5\",\"Bedford\"],[\"6\",\"Bedford\"],[\"7\",\"Bedford\"],[\"8\",\"Bedford\"],[\"9\",\"Bedford\"],[\"0\",\"Bedford Park\"],[\"1\",\"Bedford Park\"],[\"2\",\"Bedford Park\"],[\"3\",\"Bedford Park\"],[\"4\",\"Bedford Park\"],[\"5\",\"Bedford Park\"],[\"6\",\"Bedford Park\"],[\"7\",\"Bedford Park\"],[\"8\",\"Bedford Park\"],[\"9\",\"Bedford Park\"],[\"0\",\"Bellerose\"],[\"1\",\"Bellerose\"],[\"2\",\"Bellerose\"],[\"3\",\"Bellerose\"],[\"4\",\"Bellerose\"],[\"5\",\"Bellerose\"],[\"6\",\"Bellerose\"],[\"7\",\"Bellerose\"],[\"8\",\"Bellerose\"],[\"9\",\"Bellerose\"],[\"0\",\"Belmont\"],[\"1\",\"Belmont\"],[\"2\",\"Belmont\"],[\"3\",\"Belmont\"],[\"4\",\"Belmont\"],[\"5\",\"Belmont\"],[\"6\",\"Belmont\"],[\"7\",\"Belmont\"],[\"8\",\"Belmont\"],[\"9\",\"Belmont\"],[\"0\",\"Bensonhurst East\"],[\"1\",\"Bensonhurst East\"],[\"2\",\"Bensonhurst East\"],[\"3\",\"Bensonhurst East\"],[\"4\",\"Bensonhurst East\"],[\"5\",\"Bensonhurst East\"],[\"6\",\"Bensonhurst East\"],[\"7\",\"Bensonhurst East\"],[\"8\",\"Bensonhurst East\"],[\"9\",\"Bensonhurst East\"],[\"0\",\"Bensonhurst West\"],[\"1\",\"Bensonhurst West\"],[\"2\",\"Bensonhurst West\"],[\"3\",\"Bensonhurst West\"],[\"4\",\"Bensonhurst West\"],[\"5\",\"Bensonhurst West\"],[\"6\",\"Bensonhurst West\"],[\"7\",\"Bensonhurst West\"],[\"8\",\"Bensonhurst West\"],[\"9\",\"Bensonhurst West\"],[\"0\",\"Bloomfield/Emerson Hill\"],[\"1\",\"Bloomfield/Emerson Hill\"],[\"2\",\"Bloomfield/Emerson Hill\"],[\"3\",\"Bloomfield/Emerson Hill\"],[\"4\",\"Bloomfield/Emerson Hill\"],[\"5\",\"Bloomfield/Emerson Hill\"],[\"6\",\"Bloomfield/Emerson Hill\"],[\"7\",\"Bloomfield/Emerson Hill\"],[\"8\",\"Bloomfield/Emerson Hill\"],[\"9\",\"Bloomfield/Emerson Hill\"],[\"0\",\"Bloomingdale\"],[\"1\",\"Bloomingdale\"],[\"2\",\"Bloomingdale\"],[\"3\",\"Bloomingdale\"],[\"4\",\"Bloomingdale\"],[\"5\",\"Bloomingdale\"],[\"6\",\"Bloomingdale\"],[\"7\",\"Bloomingdale\"],[\"8\",\"Bloomingdale\"],[\"9\",\"Bloomingdale\"],[\"0\",\"Boerum Hill\"],[\"1\",\"Boerum Hill\"],[\"2\",\"Boerum Hill\"],[\"3\",\"Boerum Hill\"],[\"4\",\"Boerum Hill\"],[\"5\",\"Boerum Hill\"],[\"6\",\"Boerum Hill\"],[\"7\",\"Boerum Hill\"],[\"8\",\"Boerum Hill\"],[\"9\",\"Boerum Hill\"],[\"0\",\"Borough Park\"],[\"1\",\"Borough Park\"],[\"2\",\"Borough Park\"],[\"3\",\"Borough Park\"],[\"4\",\"Borough Park\"],[\"5\",\"Borough Park\"],[\"6\",\"Borough Park\"],[\"7\",\"Borough Park\"],[\"8\",\"Borough Park\"],[\"9\",\"Borough Park\"],[\"0\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"1\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"2\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"3\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"4\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"5\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"6\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"7\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"8\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"9\",\"Breezy Point/Fort Tilden/Riis Beach\"],[\"0\",\"Briarwood/Jamaica Hills\"],[\"1\",\"Briarwood/Jamaica Hills\"],[\"2\",\"Briarwood/Jamaica Hills\"],[\"3\",\"Briarwood/Jamaica Hills\"],[\"4\",\"Briarwood/Jamaica Hills\"],[\"5\",\"Briarwood/Jamaica Hills\"],[\"6\",\"Briarwood/Jamaica Hills\"],[\"7\",\"Briarwood/Jamaica Hills\"],[\"8\",\"Briarwood/Jamaica Hills\"],[\"9\",\"Briarwood/Jamaica Hills\"],[\"0\",\"Brighton Beach\"],[\"1\",\"Brighton Beach\"],[\"2\",\"Brighton Beach\"],[\"3\",\"Brighton Beach\"],[\"4\",\"Brighton Beach\"],[\"5\",\"Brighton Beach\"],[\"6\",\"Brighton Beach\"],[\"7\",\"Brighton Beach\"],[\"8\",\"Brighton Beach\"],[\"9\",\"Brighton Beach\"],[\"0\",\"Broad Channel\"],[\"1\",\"Broad Channel\"],[\"2\",\"Broad Channel\"],[\"3\",\"Broad Channel\"],[\"4\",\"Broad Channel\"],[\"5\",\"Broad Channel\"],[\"6\",\"Broad Channel\"],[\"7\",\"Broad Channel\"],[\"8\",\"Broad Channel\"],[\"9\",\"Broad Channel\"],[\"0\",\"Bronx Park\"],[\"1\",\"Bronx Park\"],[\"2\",\"Bronx Park\"],[\"3\",\"Bronx Park\"],[\"4\",\"Bronx Park\"],[\"5\",\"Bronx Park\"],[\"6\",\"Bronx Park\"],[\"7\",\"Bronx Park\"],[\"8\",\"Bronx Park\"],[\"9\",\"Bronx Park\"],[\"0\",\"Bronxdale\"],[\"1\",\"Bronxdale\"],[\"2\",\"Bronxdale\"],[\"3\",\"Bronxdale\"],[\"4\",\"Bronxdale\"],[\"5\",\"Bronxdale\"],[\"6\",\"Bronxdale\"],[\"7\",\"Bronxdale\"],[\"8\",\"Bronxdale\"],[\"9\",\"Bronxdale\"],[\"0\",\"Brooklyn Heights\"],[\"1\",\"Brooklyn Heights\"],[\"2\",\"Brooklyn Heights\"],[\"3\",\"Brooklyn Heights\"],[\"4\",\"Brooklyn Heights\"],[\"5\",\"Brooklyn Heights\"],[\"6\",\"Brooklyn Heights\"],[\"7\",\"Brooklyn Heights\"],[\"8\",\"Brooklyn Heights\"],[\"9\",\"Brooklyn Heights\"],[\"0\",\"Brooklyn Navy Yard\"],[\"1\",\"Brooklyn Navy Yard\"],[\"2\",\"Brooklyn Navy Yard\"],[\"3\",\"Brooklyn Navy Yard\"],[\"4\",\"Brooklyn Navy Yard\"],[\"5\",\"Brooklyn Navy Yard\"],[\"6\",\"Brooklyn Navy Yard\"],[\"7\",\"Brooklyn Navy Yard\"],[\"8\",\"Brooklyn Navy Yard\"],[\"9\",\"Brooklyn Navy Yard\"],[\"0\",\"Brownsville\"],[\"1\",\"Brownsville\"],[\"2\",\"Brownsville\"],[\"3\",\"Brownsville\"],[\"4\",\"Brownsville\"],[\"5\",\"Brownsville\"],[\"6\",\"Brownsville\"],[\"7\",\"Brownsville\"],[\"8\",\"Brownsville\"],[\"9\",\"Brownsville\"],[\"0\",\"Bushwick North\"],[\"1\",\"Bushwick North\"],[\"2\",\"Bushwick North\"],[\"3\",\"Bushwick North\"],[\"4\",\"Bushwick North\"],[\"5\",\"Bushwick North\"],[\"6\",\"Bushwick North\"],[\"7\",\"Bushwick North\"],[\"8\",\"Bushwick North\"],[\"9\",\"Bushwick North\"],[\"0\",\"Bushwick South\"],[\"1\",\"Bushwick South\"],[\"2\",\"Bushwick South\"],[\"3\",\"Bushwick South\"],[\"4\",\"Bushwick South\"],[\"5\",\"Bushwick South\"],[\"6\",\"Bushwick South\"],[\"7\",\"Bushwick South\"],[\"8\",\"Bushwick South\"],[\"9\",\"Bushwick South\"],[\"0\",\"Cambria Heights\"],[\"1\",\"Cambria Heights\"],[\"2\",\"Cambria Heights\"],[\"3\",\"Cambria Heights\"],[\"4\",\"Cambria Heights\"],[\"5\",\"Cambria Heights\"],[\"6\",\"Cambria Heights\"],[\"7\",\"Cambria Heights\"],[\"8\",\"Cambria Heights\"],[\"9\",\"Cambria Heights\"],[\"0\",\"Canarsie\"],[\"1\",\"Canarsie\"],[\"2\",\"Canarsie\"],[\"3\",\"Canarsie\"],[\"4\",\"Canarsie\"],[\"5\",\"Canarsie\"],[\"6\",\"Canarsie\"],[\"7\",\"Canarsie\"],[\"8\",\"Canarsie\"],[\"9\",\"Canarsie\"],[\"0\",\"Carroll Gardens\"],[\"1\",\"Carroll Gardens\"],[\"2\",\"Carroll Gardens\"],[\"3\",\"Carroll Gardens\"],[\"4\",\"Carroll Gardens\"],[\"5\",\"Carroll Gardens\"],[\"6\",\"Carroll Gardens\"],[\"7\",\"Carroll Gardens\"],[\"8\",\"Carroll Gardens\"],[\"9\",\"Carroll Gardens\"],[\"0\",\"Central Harlem\"],[\"1\",\"Central Harlem\"],[\"2\",\"Central Harlem\"],[\"3\",\"Central Harlem\"],[\"4\",\"Central Harlem\"],[\"5\",\"Central Harlem\"],[\"6\",\"Central Harlem\"],[\"7\",\"Central Harlem\"],[\"8\",\"Central Harlem\"],[\"9\",\"Central Harlem\"],[\"0\",\"Central Harlem North\"],[\"1\",\"Central Harlem North\"],[\"2\",\"Central Harlem North\"],[\"3\",\"Central Harlem North\"],[\"4\",\"Central Harlem North\"],[\"5\",\"Central Harlem North\"],[\"6\",\"Central Harlem North\"],[\"7\",\"Central Harlem North\"],[\"8\",\"Central Harlem North\"],[\"9\",\"Central Harlem North\"],[\"0\",\"Central Park\"],[\"1\",\"Central Park\"],[\"2\",\"Central Park\"],[\"3\",\"Central Park\"],[\"4\",\"Central Park\"],[\"5\",\"Central Park\"],[\"6\",\"Central Park\"],[\"7\",\"Central Park\"],[\"8\",\"Central Park\"],[\"9\",\"Central Park\"],[\"0\",\"Charleston/Tottenville\"],[\"1\",\"Charleston/Tottenville\"],[\"2\",\"Charleston/Tottenville\"],[\"3\",\"Charleston/Tottenville\"],[\"4\",\"Charleston/Tottenville\"],[\"5\",\"Charleston/Tottenville\"],[\"6\",\"Charleston/Tottenville\"],[\"7\",\"Charleston/Tottenville\"],[\"8\",\"Charleston/Tottenville\"],[\"9\",\"Charleston/Tottenville\"],[\"0\",\"Chinatown\"],[\"1\",\"Chinatown\"],[\"2\",\"Chinatown\"],[\"3\",\"Chinatown\"],[\"4\",\"Chinatown\"],[\"5\",\"Chinatown\"],[\"6\",\"Chinatown\"],[\"7\",\"Chinatown\"],[\"8\",\"Chinatown\"],[\"9\",\"Chinatown\"],[\"0\",\"City Island\"],[\"1\",\"City Island\"],[\"2\",\"City Island\"],[\"3\",\"City Island\"],[\"4\",\"City Island\"],[\"5\",\"City Island\"],[\"6\",\"City Island\"],[\"7\",\"City Island\"],[\"8\",\"City Island\"],[\"9\",\"City Island\"],[\"0\",\"Claremont/Bathgate\"],[\"1\",\"Claremont/Bathgate\"],[\"2\",\"Claremont/Bathgate\"],[\"3\",\"Claremont/Bathgate\"],[\"4\",\"Claremont/Bathgate\"],[\"5\",\"Claremont/Bathgate\"],[\"6\",\"Claremont/Bathgate\"],[\"7\",\"Claremont/Bathgate\"],[\"8\",\"Claremont/Bathgate\"],[\"9\",\"Claremont/Bathgate\"],[\"0\",\"Clinton East\"],[\"1\",\"Clinton East\"],[\"2\",\"Clinton East\"],[\"3\",\"Clinton East\"],[\"4\",\"Clinton East\"],[\"5\",\"Clinton East\"],[\"6\",\"Clinton East\"],[\"7\",\"Clinton East\"],[\"8\",\"Clinton East\"],[\"9\",\"Clinton East\"],[\"0\",\"Clinton Hill\"],[\"1\",\"Clinton Hill\"],[\"2\",\"Clinton Hill\"],[\"3\",\"Clinton Hill\"],[\"4\",\"Clinton Hill\"],[\"5\",\"Clinton Hill\"],[\"6\",\"Clinton Hill\"],[\"7\",\"Clinton Hill\"],[\"8\",\"Clinton Hill\"],[\"9\",\"Clinton Hill\"],[\"0\",\"Clinton West\"],[\"1\",\"Clinton West\"],[\"2\",\"Clinton West\"],[\"3\",\"Clinton West\"],[\"4\",\"Clinton West\"],[\"5\",\"Clinton West\"],[\"6\",\"Clinton West\"],[\"7\",\"Clinton West\"],[\"8\",\"Clinton West\"],[\"9\",\"Clinton West\"],[\"0\",\"Co-Op City\"],[\"1\",\"Co-Op City\"],[\"2\",\"Co-Op City\"],[\"3\",\"Co-Op City\"],[\"4\",\"Co-Op City\"],[\"5\",\"Co-Op City\"],[\"6\",\"Co-Op City\"],[\"7\",\"Co-Op City\"],[\"8\",\"Co-Op City\"],[\"9\",\"Co-Op City\"],[\"0\",\"Cobble Hill\"],[\"1\",\"Cobble Hill\"],[\"2\",\"Cobble Hill\"],[\"3\",\"Cobble Hill\"],[\"4\",\"Cobble Hill\"],[\"5\",\"Cobble Hill\"],[\"6\",\"Cobble Hill\"],[\"7\",\"Cobble Hill\"],[\"8\",\"Cobble Hill\"],[\"9\",\"Cobble Hill\"],[\"0\",\"College Point\"],[\"1\",\"College Point\"],[\"2\",\"College Point\"],[\"3\",\"College Point\"],[\"4\",\"College Point\"],[\"5\",\"College Point\"],[\"6\",\"College Point\"],[\"7\",\"College Point\"],[\"8\",\"College Point\"],[\"9\",\"College Point\"],[\"0\",\"Columbia Street\"],[\"1\",\"Columbia Street\"],[\"2\",\"Columbia Street\"],[\"3\",\"Columbia Street\"],[\"4\",\"Columbia Street\"],[\"5\",\"Columbia Street\"],[\"6\",\"Columbia Street\"],[\"7\",\"Columbia Street\"],[\"8\",\"Columbia Street\"],[\"9\",\"Columbia Street\"],[\"0\",\"Coney Island\"],[\"1\",\"Coney Island\"],[\"2\",\"Coney Island\"],[\"3\",\"Coney Island\"],[\"4\",\"Coney Island\"],[\"5\",\"Coney Island\"],[\"6\",\"Coney Island\"],[\"7\",\"Coney Island\"],[\"8\",\"Coney Island\"],[\"9\",\"Coney Island\"],[\"0\",\"Corona\"],[\"1\",\"Corona\"],[\"2\",\"Corona\"],[\"3\",\"Corona\"],[\"4\",\"Corona\"],[\"5\",\"Corona\"],[\"6\",\"Corona\"],[\"7\",\"Corona\"],[\"8\",\"Corona\"],[\"9\",\"Corona\"],[\"0\",\"Country Club\"],[\"1\",\"Country Club\"],[\"2\",\"Country Club\"],[\"3\",\"Country Club\"],[\"4\",\"Country Club\"],[\"5\",\"Country Club\"],[\"6\",\"Country Club\"],[\"7\",\"Country Club\"],[\"8\",\"Country Club\"],[\"9\",\"Country Club\"],[\"0\",\"Crotona Park\"],[\"1\",\"Crotona Park\"],[\"2\",\"Crotona Park\"],[\"3\",\"Crotona Park\"],[\"4\",\"Crotona Park\"],[\"5\",\"Crotona Park\"],[\"6\",\"Crotona Park\"],[\"7\",\"Crotona Park\"],[\"8\",\"Crotona Park\"],[\"9\",\"Crotona Park\"],[\"0\",\"Crotona Park East\"],[\"1\",\"Crotona Park East\"],[\"2\",\"Crotona Park East\"],[\"3\",\"Crotona Park East\"],[\"4\",\"Crotona Park East\"],[\"5\",\"Crotona Park East\"],[\"6\",\"Crotona Park East\"],[\"7\",\"Crotona Park East\"],[\"8\",\"Crotona Park East\"],[\"9\",\"Crotona Park East\"],[\"0\",\"Crown Heights North\"],[\"1\",\"Crown Heights North\"],[\"2\",\"Crown Heights North\"],[\"3\",\"Crown Heights North\"],[\"4\",\"Crown Heights North\"],[\"5\",\"Crown Heights North\"],[\"6\",\"Crown Heights North\"],[\"7\",\"Crown Heights North\"],[\"8\",\"Crown Heights North\"],[\"9\",\"Crown Heights North\"],[\"0\",\"Crown Heights South\"],[\"1\",\"Crown Heights South\"],[\"2\",\"Crown Heights South\"],[\"3\",\"Crown Heights South\"],[\"4\",\"Crown Heights South\"],[\"5\",\"Crown Heights South\"],[\"6\",\"Crown Heights South\"],[\"7\",\"Crown Heights South\"],[\"8\",\"Crown Heights South\"],[\"9\",\"Crown Heights South\"],[\"0\",\"Cypress Hills\"],[\"1\",\"Cypress Hills\"],[\"2\",\"Cypress Hills\"],[\"3\",\"Cypress Hills\"],[\"4\",\"Cypress Hills\"],[\"5\",\"Cypress Hills\"],[\"6\",\"Cypress Hills\"],[\"7\",\"Cypress Hills\"],[\"8\",\"Cypress Hills\"],[\"9\",\"Cypress Hills\"],[\"0\",\"DUMBO/Vinegar Hill\"],[\"1\",\"DUMBO/Vinegar Hill\"],[\"2\",\"DUMBO/Vinegar Hill\"],[\"3\",\"DUMBO/Vinegar Hill\"],[\"4\",\"DUMBO/Vinegar Hill\"],[\"5\",\"DUMBO/Vinegar Hill\"],[\"6\",\"DUMBO/Vinegar Hill\"],[\"7\",\"DUMBO/Vinegar Hill\"],[\"8\",\"DUMBO/Vinegar Hill\"],[\"9\",\"DUMBO/Vinegar Hill\"],[\"0\",\"Douglaston\"],[\"1\",\"Douglaston\"],[\"2\",\"Douglaston\"],[\"3\",\"Douglaston\"],[\"4\",\"Douglaston\"],[\"5\",\"Douglaston\"],[\"6\",\"Douglaston\"],[\"7\",\"Douglaston\"],[\"8\",\"Douglaston\"],[\"9\",\"Douglaston\"],[\"0\",\"Downtown Brooklyn/MetroTech\"],[\"1\",\"Downtown Brooklyn/MetroTech\"],[\"2\",\"Downtown Brooklyn/MetroTech\"],[\"3\",\"Downtown Brooklyn/MetroTech\"],[\"4\",\"Downtown Brooklyn/MetroTech\"],[\"5\",\"Downtown Brooklyn/MetroTech\"],[\"6\",\"Downtown Brooklyn/MetroTech\"],[\"7\",\"Downtown Brooklyn/MetroTech\"],[\"8\",\"Downtown Brooklyn/MetroTech\"],[\"9\",\"Downtown Brooklyn/MetroTech\"],[\"0\",\"Dyker Heights\"],[\"1\",\"Dyker Heights\"],[\"2\",\"Dyker Heights\"],[\"3\",\"Dyker Heights\"],[\"4\",\"Dyker Heights\"],[\"5\",\"Dyker Heights\"],[\"6\",\"Dyker Heights\"],[\"7\",\"Dyker Heights\"],[\"8\",\"Dyker Heights\"],[\"9\",\"Dyker Heights\"],[\"0\",\"East Chelsea\"],[\"1\",\"East Chelsea\"],[\"2\",\"East Chelsea\"],[\"3\",\"East Chelsea\"],[\"4\",\"East Chelsea\"],[\"5\",\"East Chelsea\"],[\"6\",\"East Chelsea\"],[\"7\",\"East Chelsea\"],[\"8\",\"East Chelsea\"],[\"9\",\"East Chelsea\"],[\"0\",\"East Concourse/Concourse Village\"],[\"1\",\"East Concourse/Concourse Village\"],[\"2\",\"East Concourse/Concourse Village\"],[\"3\",\"East Concourse/Concourse Village\"],[\"4\",\"East Concourse/Concourse Village\"],[\"5\",\"East Concourse/Concourse Village\"],[\"6\",\"East Concourse/Concourse Village\"],[\"7\",\"East Concourse/Concourse Village\"],[\"8\",\"East Concourse/Concourse Village\"],[\"9\",\"East Concourse/Concourse Village\"],[\"0\",\"East Elmhurst\"],[\"1\",\"East Elmhurst\"],[\"2\",\"East Elmhurst\"],[\"3\",\"East Elmhurst\"],[\"4\",\"East Elmhurst\"],[\"5\",\"East Elmhurst\"],[\"6\",\"East Elmhurst\"],[\"7\",\"East Elmhurst\"],[\"8\",\"East Elmhurst\"],[\"9\",\"East Elmhurst\"],[\"0\",\"East Flatbush/Farragut\"],[\"1\",\"East Flatbush/Farragut\"],[\"2\",\"East Flatbush/Farragut\"],[\"3\",\"East Flatbush/Farragut\"],[\"4\",\"East Flatbush/Farragut\"],[\"5\",\"East Flatbush/Farragut\"],[\"6\",\"East Flatbush/Farragut\"],[\"7\",\"East Flatbush/Farragut\"],[\"8\",\"East Flatbush/Farragut\"],[\"9\",\"East Flatbush/Farragut\"],[\"0\",\"East Flatbush/Remsen Village\"],[\"1\",\"East Flatbush/Remsen Village\"],[\"2\",\"East Flatbush/Remsen Village\"],[\"3\",\"East Flatbush/Remsen Village\"],[\"4\",\"East Flatbush/Remsen Village\"],[\"5\",\"East Flatbush/Remsen Village\"],[\"6\",\"East Flatbush/Remsen Village\"],[\"7\",\"East Flatbush/Remsen Village\"],[\"8\",\"East Flatbush/Remsen Village\"],[\"9\",\"East Flatbush/Remsen Village\"],[\"0\",\"East Flushing\"],[\"1\",\"East Flushing\"],[\"2\",\"East Flushing\"],[\"3\",\"East Flushing\"],[\"4\",\"East Flushing\"],[\"5\",\"East Flushing\"],[\"6\",\"East Flushing\"],[\"7\",\"East Flushing\"],[\"8\",\"East Flushing\"],[\"9\",\"East Flushing\"],[\"0\",\"East Harlem North\"],[\"1\",\"East Harlem North\"],[\"2\",\"East Harlem North\"],[\"3\",\"East Harlem North\"],[\"4\",\"East Harlem North\"],[\"5\",\"East Harlem North\"],[\"6\",\"East Harlem North\"],[\"7\",\"East Harlem North\"],[\"8\",\"East Harlem North\"],[\"9\",\"East Harlem North\"],[\"0\",\"East Harlem South\"],[\"1\",\"East Harlem South\"],[\"2\",\"East Harlem South\"],[\"3\",\"East Harlem South\"],[\"4\",\"East Harlem South\"],[\"5\",\"East Harlem South\"],[\"6\",\"East Harlem South\"],[\"7\",\"East Harlem South\"],[\"8\",\"East Harlem South\"],[\"9\",\"East Harlem South\"],[\"0\",\"East New York\"],[\"1\",\"East New York\"],[\"2\",\"East New York\"],[\"3\",\"East New York\"],[\"4\",\"East New York\"],[\"5\",\"East New York\"],[\"6\",\"East New York\"],[\"7\",\"East New York\"],[\"8\",\"East New York\"],[\"9\",\"East New York\"],[\"0\",\"East New York/Pennsylvania Avenue\"],[\"1\",\"East New York/Pennsylvania Avenue\"],[\"2\",\"East New York/Pennsylvania Avenue\"],[\"3\",\"East New York/Pennsylvania Avenue\"],[\"4\",\"East New York/Pennsylvania Avenue\"],[\"5\",\"East New York/Pennsylvania Avenue\"],[\"6\",\"East New York/Pennsylvania Avenue\"],[\"7\",\"East New York/Pennsylvania Avenue\"],[\"8\",\"East New York/Pennsylvania Avenue\"],[\"9\",\"East New York/Pennsylvania Avenue\"],[\"0\",\"East Tremont\"],[\"1\",\"East Tremont\"],[\"2\",\"East Tremont\"],[\"3\",\"East Tremont\"],[\"4\",\"East Tremont\"],[\"5\",\"East Tremont\"],[\"6\",\"East Tremont\"],[\"7\",\"East Tremont\"],[\"8\",\"East Tremont\"],[\"9\",\"East Tremont\"],[\"0\",\"East Village\"],[\"1\",\"East Village\"],[\"2\",\"East Village\"],[\"3\",\"East Village\"],[\"4\",\"East Village\"],[\"5\",\"East Village\"],[\"6\",\"East Village\"],[\"7\",\"East Village\"],[\"8\",\"East Village\"],[\"9\",\"East Village\"],[\"0\",\"East Williamsburg\"],[\"1\",\"East Williamsburg\"],[\"2\",\"East Williamsburg\"],[\"3\",\"East Williamsburg\"],[\"4\",\"East Williamsburg\"],[\"5\",\"East Williamsburg\"],[\"6\",\"East Williamsburg\"],[\"7\",\"East Williamsburg\"],[\"8\",\"East Williamsburg\"],[\"9\",\"East Williamsburg\"],[\"0\",\"Eastchester\"],[\"1\",\"Eastchester\"],[\"2\",\"Eastchester\"],[\"3\",\"Eastchester\"],[\"4\",\"Eastchester\"],[\"5\",\"Eastchester\"],[\"6\",\"Eastchester\"],[\"7\",\"Eastchester\"],[\"8\",\"Eastchester\"],[\"9\",\"Eastchester\"],[\"0\",\"Elmhurst\"],[\"1\",\"Elmhurst\"],[\"2\",\"Elmhurst\"],[\"3\",\"Elmhurst\"],[\"4\",\"Elmhurst\"],[\"5\",\"Elmhurst\"],[\"6\",\"Elmhurst\"],[\"7\",\"Elmhurst\"],[\"8\",\"Elmhurst\"],[\"9\",\"Elmhurst\"],[\"0\",\"Elmhurst/Maspeth\"],[\"1\",\"Elmhurst/Maspeth\"],[\"2\",\"Elmhurst/Maspeth\"],[\"3\",\"Elmhurst/Maspeth\"],[\"4\",\"Elmhurst/Maspeth\"],[\"5\",\"Elmhurst/Maspeth\"],[\"6\",\"Elmhurst/Maspeth\"],[\"7\",\"Elmhurst/Maspeth\"],[\"8\",\"Elmhurst/Maspeth\"],[\"9\",\"Elmhurst/Maspeth\"],[\"0\",\"Eltingville/Annadale/Prince's Bay\"],[\"1\",\"Eltingville/Annadale/Prince's Bay\"],[\"2\",\"Eltingville/Annadale/Prince's Bay\"],[\"3\",\"Eltingville/Annadale/Prince's Bay\"],[\"4\",\"Eltingville/Annadale/Prince's Bay\"],[\"5\",\"Eltingville/Annadale/Prince's Bay\"],[\"6\",\"Eltingville/Annadale/Prince's Bay\"],[\"7\",\"Eltingville/Annadale/Prince's Bay\"],[\"8\",\"Eltingville/Annadale/Prince's Bay\"],[\"9\",\"Eltingville/Annadale/Prince's Bay\"],[\"0\",\"Erasmus\"],[\"1\",\"Erasmus\"],[\"2\",\"Erasmus\"],[\"3\",\"Erasmus\"],[\"4\",\"Erasmus\"],[\"5\",\"Erasmus\"],[\"6\",\"Erasmus\"],[\"7\",\"Erasmus\"],[\"8\",\"Erasmus\"],[\"9\",\"Erasmus\"],[\"0\",\"Far Rockaway\"],[\"1\",\"Far Rockaway\"],[\"2\",\"Far Rockaway\"],[\"3\",\"Far Rockaway\"],[\"4\",\"Far Rockaway\"],[\"5\",\"Far Rockaway\"],[\"6\",\"Far Rockaway\"],[\"7\",\"Far Rockaway\"],[\"8\",\"Far Rockaway\"],[\"9\",\"Far Rockaway\"],[\"0\",\"Financial District North\"],[\"1\",\"Financial District North\"],[\"2\",\"Financial District North\"],[\"3\",\"Financial District North\"],[\"4\",\"Financial District North\"],[\"5\",\"Financial District North\"],[\"6\",\"Financial District North\"],[\"7\",\"Financial District North\"],[\"8\",\"Financial District North\"],[\"9\",\"Financial District North\"],[\"0\",\"Financial District South\"],[\"1\",\"Financial District South\"],[\"2\",\"Financial District South\"],[\"3\",\"Financial District South\"],[\"4\",\"Financial District South\"],[\"5\",\"Financial District South\"],[\"6\",\"Financial District South\"],[\"7\",\"Financial District South\"],[\"8\",\"Financial District South\"],[\"9\",\"Financial District South\"],[\"0\",\"Flatbush/Ditmas Park\"],[\"1\",\"Flatbush/Ditmas Park\"],[\"2\",\"Flatbush/Ditmas Park\"],[\"3\",\"Flatbush/Ditmas Park\"],[\"4\",\"Flatbush/Ditmas Park\"],[\"5\",\"Flatbush/Ditmas Park\"],[\"6\",\"Flatbush/Ditmas Park\"],[\"7\",\"Flatbush/Ditmas Park\"],[\"8\",\"Flatbush/Ditmas Park\"],[\"9\",\"Flatbush/Ditmas Park\"],[\"0\",\"Flatiron\"],[\"1\",\"Flatiron\"],[\"2\",\"Flatiron\"],[\"3\",\"Flatiron\"],[\"4\",\"Flatiron\"],[\"5\",\"Flatiron\"],[\"6\",\"Flatiron\"],[\"7\",\"Flatiron\"],[\"8\",\"Flatiron\"],[\"9\",\"Flatiron\"],[\"0\",\"Flatlands\"],[\"1\",\"Flatlands\"],[\"2\",\"Flatlands\"],[\"3\",\"Flatlands\"],[\"4\",\"Flatlands\"],[\"5\",\"Flatlands\"],[\"6\",\"Flatlands\"],[\"7\",\"Flatlands\"],[\"8\",\"Flatlands\"],[\"9\",\"Flatlands\"],[\"0\",\"Flushing\"],[\"1\",\"Flushing\"],[\"2\",\"Flushing\"],[\"3\",\"Flushing\"],[\"4\",\"Flushing\"],[\"5\",\"Flushing\"],[\"6\",\"Flushing\"],[\"7\",\"Flushing\"],[\"8\",\"Flushing\"],[\"9\",\"Flushing\"],[\"0\",\"Flushing Meadows-Corona Park\"],[\"1\",\"Flushing Meadows-Corona Park\"],[\"2\",\"Flushing Meadows-Corona Park\"],[\"3\",\"Flushing Meadows-Corona Park\"],[\"4\",\"Flushing Meadows-Corona Park\"],[\"5\",\"Flushing Meadows-Corona Park\"],[\"6\",\"Flushing Meadows-Corona Park\"],[\"7\",\"Flushing Meadows-Corona Park\"],[\"8\",\"Flushing Meadows-Corona Park\"],[\"9\",\"Flushing Meadows-Corona Park\"],[\"0\",\"Fordham South\"],[\"1\",\"Fordham South\"],[\"2\",\"Fordham South\"],[\"3\",\"Fordham South\"],[\"4\",\"Fordham South\"],[\"5\",\"Fordham South\"],[\"6\",\"Fordham South\"],[\"7\",\"Fordham South\"],[\"8\",\"Fordham South\"],[\"9\",\"Fordham South\"],[\"0\",\"Forest Hills\"],[\"1\",\"Forest Hills\"],[\"2\",\"Forest Hills\"],[\"3\",\"Forest Hills\"],[\"4\",\"Forest Hills\"],[\"5\",\"Forest Hills\"],[\"6\",\"Forest Hills\"],[\"7\",\"Forest Hills\"],[\"8\",\"Forest Hills\"],[\"9\",\"Forest Hills\"],[\"0\",\"Forest Park/Highland Park\"],[\"1\",\"Forest Park/Highland Park\"],[\"2\",\"Forest Park/Highland Park\"],[\"3\",\"Forest Park/Highland Park\"],[\"4\",\"Forest Park/Highland Park\"],[\"5\",\"Forest Park/Highland Park\"],[\"6\",\"Forest Park/Highland Park\"],[\"7\",\"Forest Park/Highland Park\"],[\"8\",\"Forest Park/Highland Park\"],[\"9\",\"Forest Park/Highland Park\"],[\"0\",\"Fort Greene\"],[\"1\",\"Fort Greene\"],[\"2\",\"Fort Greene\"],[\"3\",\"Fort Greene\"],[\"4\",\"Fort Greene\"],[\"5\",\"Fort Greene\"],[\"6\",\"Fort Greene\"],[\"7\",\"Fort Greene\"],[\"8\",\"Fort Greene\"],[\"9\",\"Fort Greene\"],[\"0\",\"Fresh Meadows\"],[\"1\",\"Fresh Meadows\"],[\"2\",\"Fresh Meadows\"],[\"3\",\"Fresh Meadows\"],[\"4\",\"Fresh Meadows\"],[\"5\",\"Fresh Meadows\"],[\"6\",\"Fresh Meadows\"],[\"7\",\"Fresh Meadows\"],[\"8\",\"Fresh Meadows\"],[\"9\",\"Fresh Meadows\"],[\"0\",\"Freshkills Park\"],[\"1\",\"Freshkills Park\"],[\"2\",\"Freshkills Park\"],[\"3\",\"Freshkills Park\"],[\"4\",\"Freshkills Park\"],[\"5\",\"Freshkills Park\"],[\"6\",\"Freshkills Park\"],[\"7\",\"Freshkills Park\"],[\"8\",\"Freshkills Park\"],[\"9\",\"Freshkills Park\"],[\"0\",\"Garment District\"],[\"1\",\"Garment District\"],[\"2\",\"Garment District\"],[\"3\",\"Garment District\"],[\"4\",\"Garment District\"],[\"5\",\"Garment District\"],[\"6\",\"Garment District\"],[\"7\",\"Garment District\"],[\"8\",\"Garment District\"],[\"9\",\"Garment District\"],[\"0\",\"Glen Oaks\"],[\"1\",\"Glen Oaks\"],[\"2\",\"Glen Oaks\"],[\"3\",\"Glen Oaks\"],[\"4\",\"Glen Oaks\"],[\"5\",\"Glen Oaks\"],[\"6\",\"Glen Oaks\"],[\"7\",\"Glen Oaks\"],[\"8\",\"Glen Oaks\"],[\"9\",\"Glen Oaks\"],[\"0\",\"Glendale\"],[\"1\",\"Glendale\"],[\"2\",\"Glendale\"],[\"3\",\"Glendale\"],[\"4\",\"Glendale\"],[\"5\",\"Glendale\"],[\"6\",\"Glendale\"],[\"7\",\"Glendale\"],[\"8\",\"Glendale\"],[\"9\",\"Glendale\"],[\"0\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"1\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"2\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"3\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"4\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"5\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"6\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"7\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"8\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"9\",\"Governor's Island/Ellis Island/Liberty Island\"],[\"0\",\"Gowanus\"],[\"1\",\"Gowanus\"],[\"2\",\"Gowanus\"],[\"3\",\"Gowanus\"],[\"4\",\"Gowanus\"],[\"5\",\"Gowanus\"],[\"6\",\"Gowanus\"],[\"7\",\"Gowanus\"],[\"8\",\"Gowanus\"],[\"9\",\"Gowanus\"],[\"0\",\"Gramercy\"],[\"1\",\"Gramercy\"],[\"2\",\"Gramercy\"],[\"3\",\"Gramercy\"],[\"4\",\"Gramercy\"],[\"5\",\"Gramercy\"],[\"6\",\"Gramercy\"],[\"7\",\"Gramercy\"],[\"8\",\"Gramercy\"],[\"9\",\"Gramercy\"],[\"0\",\"Gravesend\"],[\"1\",\"Gravesend\"],[\"2\",\"Gravesend\"],[\"3\",\"Gravesend\"],[\"4\",\"Gravesend\"],[\"5\",\"Gravesend\"],[\"6\",\"Gravesend\"],[\"7\",\"Gravesend\"],[\"8\",\"Gravesend\"],[\"9\",\"Gravesend\"],[\"0\",\"Great Kills\"],[\"1\",\"Great Kills\"],[\"2\",\"Great Kills\"],[\"3\",\"Great Kills\"],[\"4\",\"Great Kills\"],[\"5\",\"Great Kills\"],[\"6\",\"Great Kills\"],[\"7\",\"Great Kills\"],[\"8\",\"Great Kills\"],[\"9\",\"Great Kills\"],[\"0\",\"Great Kills Park\"],[\"1\",\"Great Kills Park\"],[\"2\",\"Great Kills Park\"],[\"3\",\"Great Kills Park\"],[\"4\",\"Great Kills Park\"],[\"5\",\"Great Kills Park\"],[\"6\",\"Great Kills Park\"],[\"7\",\"Great Kills Park\"],[\"8\",\"Great Kills Park\"],[\"9\",\"Great Kills Park\"],[\"0\",\"Green-Wood Cemetery\"],[\"1\",\"Green-Wood Cemetery\"],[\"2\",\"Green-Wood Cemetery\"],[\"3\",\"Green-Wood Cemetery\"],[\"4\",\"Green-Wood Cemetery\"],[\"5\",\"Green-Wood Cemetery\"],[\"6\",\"Green-Wood Cemetery\"],[\"7\",\"Green-Wood Cemetery\"],[\"8\",\"Green-Wood Cemetery\"],[\"9\",\"Green-Wood Cemetery\"],[\"0\",\"Greenpoint\"],[\"1\",\"Greenpoint\"],[\"2\",\"Greenpoint\"],[\"3\",\"Greenpoint\"],[\"4\",\"Greenpoint\"],[\"5\",\"Greenpoint\"],[\"6\",\"Greenpoint\"],[\"7\",\"Greenpoint\"],[\"8\",\"Greenpoint\"],[\"9\",\"Greenpoint\"],[\"0\",\"Greenwich Village North\"],[\"1\",\"Greenwich Village North\"],[\"2\",\"Greenwich Village North\"],[\"3\",\"Greenwich Village North\"],[\"4\",\"Greenwich Village North\"],[\"5\",\"Greenwich Village North\"],[\"6\",\"Greenwich Village North\"],[\"7\",\"Greenwich Village North\"],[\"8\",\"Greenwich Village North\"],[\"9\",\"Greenwich Village North\"],[\"0\",\"Greenwich Village South\"],[\"1\",\"Greenwich Village South\"],[\"2\",\"Greenwich Village South\"],[\"3\",\"Greenwich Village South\"],[\"4\",\"Greenwich Village South\"],[\"5\",\"Greenwich Village South\"],[\"6\",\"Greenwich Village South\"],[\"7\",\"Greenwich Village South\"],[\"8\",\"Greenwich Village South\"],[\"9\",\"Greenwich Village South\"],[\"0\",\"Grymes Hill/Clifton\"],[\"1\",\"Grymes Hill/Clifton\"],[\"2\",\"Grymes Hill/Clifton\"],[\"3\",\"Grymes Hill/Clifton\"],[\"4\",\"Grymes Hill/Clifton\"],[\"5\",\"Grymes Hill/Clifton\"],[\"6\",\"Grymes Hill/Clifton\"],[\"7\",\"Grymes Hill/Clifton\"],[\"8\",\"Grymes Hill/Clifton\"],[\"9\",\"Grymes Hill/Clifton\"],[\"0\",\"Hamilton Heights\"],[\"1\",\"Hamilton Heights\"],[\"2\",\"Hamilton Heights\"],[\"3\",\"Hamilton Heights\"],[\"4\",\"Hamilton Heights\"],[\"5\",\"Hamilton Heights\"],[\"6\",\"Hamilton Heights\"],[\"7\",\"Hamilton Heights\"],[\"8\",\"Hamilton Heights\"],[\"9\",\"Hamilton Heights\"],[\"0\",\"Hammels/Arverne\"],[\"1\",\"Hammels/Arverne\"],[\"2\",\"Hammels/Arverne\"],[\"3\",\"Hammels/Arverne\"],[\"4\",\"Hammels/Arverne\"],[\"5\",\"Hammels/Arverne\"],[\"6\",\"Hammels/Arverne\"],[\"7\",\"Hammels/Arverne\"],[\"8\",\"Hammels/Arverne\"],[\"9\",\"Hammels/Arverne\"],[\"0\",\"Heartland Village/Todt Hill\"],[\"1\",\"Heartland Village/Todt Hill\"],[\"2\",\"Heartland Village/Todt Hill\"],[\"3\",\"Heartland Village/Todt Hill\"],[\"4\",\"Heartland Village/Todt Hill\"],[\"5\",\"Heartland Village/Todt Hill\"],[\"6\",\"Heartland Village/Todt Hill\"],[\"7\",\"Heartland Village/Todt Hill\"],[\"8\",\"Heartland Village/Todt Hill\"],[\"9\",\"Heartland Village/Todt Hill\"],[\"0\",\"Highbridge\"],[\"1\",\"Highbridge\"],[\"2\",\"Highbridge\"],[\"3\",\"Highbridge\"],[\"4\",\"Highbridge\"],[\"5\",\"Highbridge\"],[\"6\",\"Highbridge\"],[\"7\",\"Highbridge\"],[\"8\",\"Highbridge\"],[\"9\",\"Highbridge\"],[\"0\",\"Highbridge Park\"],[\"1\",\"Highbridge Park\"],[\"2\",\"Highbridge Park\"],[\"3\",\"Highbridge Park\"],[\"4\",\"Highbridge Park\"],[\"5\",\"Highbridge Park\"],[\"6\",\"Highbridge Park\"],[\"7\",\"Highbridge Park\"],[\"8\",\"Highbridge Park\"],[\"9\",\"Highbridge Park\"],[\"0\",\"Hillcrest/Pomonok\"],[\"1\",\"Hillcrest/Pomonok\"],[\"2\",\"Hillcrest/Pomonok\"],[\"3\",\"Hillcrest/Pomonok\"],[\"4\",\"Hillcrest/Pomonok\"],[\"5\",\"Hillcrest/Pomonok\"],[\"6\",\"Hillcrest/Pomonok\"],[\"7\",\"Hillcrest/Pomonok\"],[\"8\",\"Hillcrest/Pomonok\"],[\"9\",\"Hillcrest/Pomonok\"],[\"0\",\"Hollis\"],[\"1\",\"Hollis\"],[\"2\",\"Hollis\"],[\"3\",\"Hollis\"],[\"4\",\"Hollis\"],[\"5\",\"Hollis\"],[\"6\",\"Hollis\"],[\"7\",\"Hollis\"],[\"8\",\"Hollis\"],[\"9\",\"Hollis\"],[\"0\",\"Homecrest\"],[\"1\",\"Homecrest\"],[\"2\",\"Homecrest\"],[\"3\",\"Homecrest\"],[\"4\",\"Homecrest\"],[\"5\",\"Homecrest\"],[\"6\",\"Homecrest\"],[\"7\",\"Homecrest\"],[\"8\",\"Homecrest\"],[\"9\",\"Homecrest\"],[\"0\",\"Howard Beach\"],[\"1\",\"Howard Beach\"],[\"2\",\"Howard Beach\"],[\"3\",\"Howard Beach\"],[\"4\",\"Howard Beach\"],[\"5\",\"Howard Beach\"],[\"6\",\"Howard Beach\"],[\"7\",\"Howard Beach\"],[\"8\",\"Howard Beach\"],[\"9\",\"Howard Beach\"],[\"0\",\"Hudson Sq\"],[\"1\",\"Hudson Sq\"],[\"2\",\"Hudson Sq\"],[\"3\",\"Hudson Sq\"],[\"4\",\"Hudson Sq\"],[\"5\",\"Hudson Sq\"],[\"6\",\"Hudson Sq\"],[\"7\",\"Hudson Sq\"],[\"8\",\"Hudson Sq\"],[\"9\",\"Hudson Sq\"],[\"0\",\"Hunts Point\"],[\"1\",\"Hunts Point\"],[\"2\",\"Hunts Point\"],[\"3\",\"Hunts Point\"],[\"4\",\"Hunts Point\"],[\"5\",\"Hunts Point\"],[\"6\",\"Hunts Point\"],[\"7\",\"Hunts Point\"],[\"8\",\"Hunts Point\"],[\"9\",\"Hunts Point\"],[\"0\",\"Inwood\"],[\"1\",\"Inwood\"],[\"2\",\"Inwood\"],[\"3\",\"Inwood\"],[\"4\",\"Inwood\"],[\"5\",\"Inwood\"],[\"6\",\"Inwood\"],[\"7\",\"Inwood\"],[\"8\",\"Inwood\"],[\"9\",\"Inwood\"],[\"0\",\"Inwood Hill Park\"],[\"1\",\"Inwood Hill Park\"],[\"2\",\"Inwood Hill Park\"],[\"3\",\"Inwood Hill Park\"],[\"4\",\"Inwood Hill Park\"],[\"5\",\"Inwood Hill Park\"],[\"6\",\"Inwood Hill Park\"],[\"7\",\"Inwood Hill Park\"],[\"8\",\"Inwood Hill Park\"],[\"9\",\"Inwood Hill Park\"],[\"0\",\"JFK Airport\"],[\"1\",\"JFK Airport\"],[\"2\",\"JFK Airport\"],[\"3\",\"JFK Airport\"],[\"4\",\"JFK Airport\"],[\"5\",\"JFK Airport\"],[\"6\",\"JFK Airport\"],[\"7\",\"JFK Airport\"],[\"8\",\"JFK Airport\"],[\"9\",\"JFK Airport\"],[\"0\",\"Jackson Heights\"],[\"1\",\"Jackson Heights\"],[\"2\",\"Jackson Heights\"],[\"3\",\"Jackson Heights\"],[\"4\",\"Jackson Heights\"],[\"5\",\"Jackson Heights\"],[\"6\",\"Jackson Heights\"],[\"7\",\"Jackson Heights\"],[\"8\",\"Jackson Heights\"],[\"9\",\"Jackson Heights\"],[\"0\",\"Jamaica\"],[\"1\",\"Jamaica\"],[\"2\",\"Jamaica\"],[\"3\",\"Jamaica\"],[\"4\",\"Jamaica\"],[\"5\",\"Jamaica\"],[\"6\",\"Jamaica\"],[\"7\",\"Jamaica\"],[\"8\",\"Jamaica\"],[\"9\",\"Jamaica\"],[\"0\",\"Jamaica Bay\"],[\"1\",\"Jamaica Bay\"],[\"2\",\"Jamaica Bay\"],[\"3\",\"Jamaica Bay\"],[\"4\",\"Jamaica Bay\"],[\"5\",\"Jamaica Bay\"],[\"6\",\"Jamaica Bay\"],[\"7\",\"Jamaica Bay\"],[\"8\",\"Jamaica Bay\"],[\"9\",\"Jamaica Bay\"],[\"0\",\"Jamaica Estates\"],[\"1\",\"Jamaica Estates\"],[\"2\",\"Jamaica Estates\"],[\"3\",\"Jamaica Estates\"],[\"4\",\"Jamaica Estates\"],[\"5\",\"Jamaica Estates\"],[\"6\",\"Jamaica Estates\"],[\"7\",\"Jamaica Estates\"],[\"8\",\"Jamaica Estates\"],[\"9\",\"Jamaica Estates\"],[\"0\",\"Kensington\"],[\"1\",\"Kensington\"],[\"2\",\"Kensington\"],[\"3\",\"Kensington\"],[\"4\",\"Kensington\"],[\"5\",\"Kensington\"],[\"6\",\"Kensington\"],[\"7\",\"Kensington\"],[\"8\",\"Kensington\"],[\"9\",\"Kensington\"],[\"0\",\"Kew Gardens\"],[\"1\",\"Kew Gardens\"],[\"2\",\"Kew Gardens\"],[\"3\",\"Kew Gardens\"],[\"4\",\"Kew Gardens\"],[\"5\",\"Kew Gardens\"],[\"6\",\"Kew Gardens\"],[\"7\",\"Kew Gardens\"],[\"8\",\"Kew Gardens\"],[\"9\",\"Kew Gardens\"],[\"0\",\"Kew Gardens Hills\"],[\"1\",\"Kew Gardens Hills\"],[\"2\",\"Kew Gardens Hills\"],[\"3\",\"Kew Gardens Hills\"],[\"4\",\"Kew Gardens Hills\"],[\"5\",\"Kew Gardens Hills\"],[\"6\",\"Kew Gardens Hills\"],[\"7\",\"Kew Gardens Hills\"],[\"8\",\"Kew Gardens Hills\"],[\"9\",\"Kew Gardens Hills\"],[\"0\",\"Kingsbridge Heights\"],[\"1\",\"Kingsbridge Heights\"],[\"2\",\"Kingsbridge Heights\"],[\"3\",\"Kingsbridge Heights\"],[\"4\",\"Kingsbridge Heights\"],[\"5\",\"Kingsbridge Heights\"],[\"6\",\"Kingsbridge Heights\"],[\"7\",\"Kingsbridge Heights\"],[\"8\",\"Kingsbridge Heights\"],[\"9\",\"Kingsbridge Heights\"],[\"0\",\"Kips Bay\"],[\"1\",\"Kips Bay\"],[\"2\",\"Kips Bay\"],[\"3\",\"Kips Bay\"],[\"4\",\"Kips Bay\"],[\"5\",\"Kips Bay\"],[\"6\",\"Kips Bay\"],[\"7\",\"Kips Bay\"],[\"8\",\"Kips Bay\"],[\"9\",\"Kips Bay\"],[\"0\",\"LaGuardia Airport\"],[\"1\",\"LaGuardia Airport\"],[\"2\",\"LaGuardia Airport\"],[\"3\",\"LaGuardia Airport\"],[\"4\",\"LaGuardia Airport\"],[\"5\",\"LaGuardia Airport\"],[\"6\",\"LaGuardia Airport\"],[\"7\",\"LaGuardia Airport\"],[\"8\",\"LaGuardia Airport\"],[\"9\",\"LaGuardia Airport\"],[\"0\",\"Laurelton\"],[\"1\",\"Laurelton\"],[\"2\",\"Laurelton\"],[\"3\",\"Laurelton\"],[\"4\",\"Laurelton\"],[\"5\",\"Laurelton\"],[\"6\",\"Laurelton\"],[\"7\",\"Laurelton\"],[\"8\",\"Laurelton\"],[\"9\",\"Laurelton\"],[\"0\",\"Lenox Hill East\"],[\"1\",\"Lenox Hill East\"],[\"2\",\"Lenox Hill East\"],[\"3\",\"Lenox Hill East\"],[\"4\",\"Lenox Hill East\"],[\"5\",\"Lenox Hill East\"],[\"6\",\"Lenox Hill East\"],[\"7\",\"Lenox Hill East\"],[\"8\",\"Lenox Hill East\"],[\"9\",\"Lenox Hill East\"],[\"0\",\"Lenox Hill West\"],[\"1\",\"Lenox Hill West\"],[\"2\",\"Lenox Hill West\"],[\"3\",\"Lenox Hill West\"],[\"4\",\"Lenox Hill West\"],[\"5\",\"Lenox Hill West\"],[\"6\",\"Lenox Hill West\"],[\"7\",\"Lenox Hill West\"],[\"8\",\"Lenox Hill West\"],[\"9\",\"Lenox Hill West\"],[\"0\",\"Lincoln Square East\"],[\"1\",\"Lincoln Square East\"],[\"2\",\"Lincoln Square East\"],[\"3\",\"Lincoln Square East\"],[\"4\",\"Lincoln Square East\"],[\"5\",\"Lincoln Square East\"],[\"6\",\"Lincoln Square East\"],[\"7\",\"Lincoln Square East\"],[\"8\",\"Lincoln Square East\"],[\"9\",\"Lincoln Square East\"],[\"0\",\"Lincoln Square West\"],[\"1\",\"Lincoln Square West\"],[\"2\",\"Lincoln Square West\"],[\"3\",\"Lincoln Square West\"],[\"4\",\"Lincoln Square West\"],[\"5\",\"Lincoln Square West\"],[\"6\",\"Lincoln Square West\"],[\"7\",\"Lincoln Square West\"],[\"8\",\"Lincoln Square West\"],[\"9\",\"Lincoln Square West\"],[\"0\",\"Little Italy/NoLiTa\"],[\"1\",\"Little Italy/NoLiTa\"],[\"2\",\"Little Italy/NoLiTa\"],[\"3\",\"Little Italy/NoLiTa\"],[\"4\",\"Little Italy/NoLiTa\"],[\"5\",\"Little Italy/NoLiTa\"],[\"6\",\"Little Italy/NoLiTa\"],[\"7\",\"Little Italy/NoLiTa\"],[\"8\",\"Little Italy/NoLiTa\"],[\"9\",\"Little Italy/NoLiTa\"],[\"0\",\"Long Island City/Hunters Point\"],[\"1\",\"Long Island City/Hunters Point\"],[\"2\",\"Long Island City/Hunters Point\"],[\"3\",\"Long Island City/Hunters Point\"],[\"4\",\"Long Island City/Hunters Point\"],[\"5\",\"Long Island City/Hunters Point\"],[\"6\",\"Long Island City/Hunters Point\"],[\"7\",\"Long Island City/Hunters Point\"],[\"8\",\"Long Island City/Hunters Point\"],[\"9\",\"Long Island City/Hunters Point\"],[\"0\",\"Long Island City/Queens Plaza\"],[\"1\",\"Long Island City/Queens Plaza\"],[\"2\",\"Long Island City/Queens Plaza\"],[\"3\",\"Long Island City/Queens Plaza\"],[\"4\",\"Long Island City/Queens Plaza\"],[\"5\",\"Long Island City/Queens Plaza\"],[\"6\",\"Long Island City/Queens Plaza\"],[\"7\",\"Long Island City/Queens Plaza\"],[\"8\",\"Long Island City/Queens Plaza\"],[\"9\",\"Long Island City/Queens Plaza\"],[\"0\",\"Longwood\"],[\"1\",\"Longwood\"],[\"2\",\"Longwood\"],[\"3\",\"Longwood\"],[\"4\",\"Longwood\"],[\"5\",\"Longwood\"],[\"6\",\"Longwood\"],[\"7\",\"Longwood\"],[\"8\",\"Longwood\"],[\"9\",\"Longwood\"],[\"0\",\"Lower East Side\"],[\"1\",\"Lower East Side\"],[\"2\",\"Lower East Side\"],[\"3\",\"Lower East Side\"],[\"4\",\"Lower East Side\"],[\"5\",\"Lower East Side\"],[\"6\",\"Lower East Side\"],[\"7\",\"Lower East Side\"],[\"8\",\"Lower East Side\"],[\"9\",\"Lower East Side\"],[\"0\",\"Madison\"],[\"1\",\"Madison\"],[\"2\",\"Madison\"],[\"3\",\"Madison\"],[\"4\",\"Madison\"],[\"5\",\"Madison\"],[\"6\",\"Madison\"],[\"7\",\"Madison\"],[\"8\",\"Madison\"],[\"9\",\"Madison\"],[\"0\",\"Manhattan Beach\"],[\"1\",\"Manhattan Beach\"],[\"2\",\"Manhattan Beach\"],[\"3\",\"Manhattan Beach\"],[\"4\",\"Manhattan Beach\"],[\"5\",\"Manhattan Beach\"],[\"6\",\"Manhattan Beach\"],[\"7\",\"Manhattan Beach\"],[\"8\",\"Manhattan Beach\"],[\"9\",\"Manhattan Beach\"],[\"0\",\"Manhattan Valley\"],[\"1\",\"Manhattan Valley\"],[\"2\",\"Manhattan Valley\"],[\"3\",\"Manhattan Valley\"],[\"4\",\"Manhattan Valley\"],[\"5\",\"Manhattan Valley\"],[\"6\",\"Manhattan Valley\"],[\"7\",\"Manhattan Valley\"],[\"8\",\"Manhattan Valley\"],[\"9\",\"Manhattan Valley\"],[\"0\",\"Manhattanville\"],[\"1\",\"Manhattanville\"],[\"2\",\"Manhattanville\"],[\"3\",\"Manhattanville\"],[\"4\",\"Manhattanville\"],[\"5\",\"Manhattanville\"],[\"6\",\"Manhattanville\"],[\"7\",\"Manhattanville\"],[\"8\",\"Manhattanville\"],[\"9\",\"Manhattanville\"],[\"0\",\"Marble Hill\"],[\"1\",\"Marble Hill\"],[\"2\",\"Marble Hill\"],[\"3\",\"Marble Hill\"],[\"4\",\"Marble Hill\"],[\"5\",\"Marble Hill\"],[\"6\",\"Marble Hill\"],[\"7\",\"Marble Hill\"],[\"8\",\"Marble Hill\"],[\"9\",\"Marble Hill\"],[\"0\",\"Marine Park/Floyd Bennett Field\"],[\"1\",\"Marine Park/Floyd Bennett Field\"],[\"2\",\"Marine Park/Floyd Bennett Field\"],[\"3\",\"Marine Park/Floyd Bennett Field\"],[\"4\",\"Marine Park/Floyd Bennett Field\"],[\"5\",\"Marine Park/Floyd Bennett Field\"],[\"6\",\"Marine Park/Floyd Bennett Field\"],[\"7\",\"Marine Park/Floyd Bennett Field\"],[\"8\",\"Marine Park/Floyd Bennett Field\"],[\"9\",\"Marine Park/Floyd Bennett Field\"],[\"0\",\"Marine Park/Mill Basin\"],[\"1\",\"Marine Park/Mill Basin\"],[\"2\",\"Marine Park/Mill Basin\"],[\"3\",\"Marine Park/Mill Basin\"],[\"4\",\"Marine Park/Mill Basin\"],[\"5\",\"Marine Park/Mill Basin\"],[\"6\",\"Marine Park/Mill Basin\"],[\"7\",\"Marine Park/Mill Basin\"],[\"8\",\"Marine Park/Mill Basin\"],[\"9\",\"Marine Park/Mill Basin\"],[\"0\",\"Mariners Harbor\"],[\"1\",\"Mariners Harbor\"],[\"2\",\"Mariners Harbor\"],[\"3\",\"Mariners Harbor\"],[\"4\",\"Mariners Harbor\"],[\"5\",\"Mariners Harbor\"],[\"6\",\"Mariners Harbor\"],[\"7\",\"Mariners Harbor\"],[\"8\",\"Mariners Harbor\"],[\"9\",\"Mariners Harbor\"],[\"0\",\"Maspeth\"],[\"1\",\"Maspeth\"],[\"2\",\"Maspeth\"],[\"3\",\"Maspeth\"],[\"4\",\"Maspeth\"],[\"5\",\"Maspeth\"],[\"6\",\"Maspeth\"],[\"7\",\"Maspeth\"],[\"8\",\"Maspeth\"],[\"9\",\"Maspeth\"],[\"0\",\"Meatpacking/West Village West\"],[\"1\",\"Meatpacking/West Village West\"],[\"2\",\"Meatpacking/West Village West\"],[\"3\",\"Meatpacking/West Village West\"],[\"4\",\"Meatpacking/West Village West\"],[\"5\",\"Meatpacking/West Village West\"],[\"6\",\"Meatpacking/West Village West\"],[\"7\",\"Meatpacking/West Village West\"],[\"8\",\"Meatpacking/West Village West\"],[\"9\",\"Meatpacking/West Village West\"],[\"0\",\"Melrose South\"],[\"1\",\"Melrose South\"],[\"2\",\"Melrose South\"],[\"3\",\"Melrose South\"],[\"4\",\"Melrose South\"],[\"5\",\"Melrose South\"],[\"6\",\"Melrose South\"],[\"7\",\"Melrose South\"],[\"8\",\"Melrose South\"],[\"9\",\"Melrose South\"],[\"0\",\"Middle Village\"],[\"1\",\"Middle Village\"],[\"2\",\"Middle Village\"],[\"3\",\"Middle Village\"],[\"4\",\"Middle Village\"],[\"5\",\"Middle Village\"],[\"6\",\"Middle Village\"],[\"7\",\"Middle Village\"],[\"8\",\"Middle Village\"],[\"9\",\"Middle Village\"],[\"0\",\"Midtown Center\"],[\"1\",\"Midtown Center\"],[\"2\",\"Midtown Center\"],[\"3\",\"Midtown Center\"],[\"4\",\"Midtown Center\"],[\"5\",\"Midtown Center\"],[\"6\",\"Midtown Center\"],[\"7\",\"Midtown Center\"],[\"8\",\"Midtown Center\"],[\"9\",\"Midtown Center\"],[\"0\",\"Midtown East\"],[\"1\",\"Midtown East\"],[\"2\",\"Midtown East\"],[\"3\",\"Midtown East\"],[\"4\",\"Midtown East\"],[\"5\",\"Midtown East\"],[\"6\",\"Midtown East\"],[\"7\",\"Midtown East\"],[\"8\",\"Midtown East\"],[\"9\",\"Midtown East\"],[\"0\",\"Midtown North\"],[\"1\",\"Midtown North\"],[\"2\",\"Midtown North\"],[\"3\",\"Midtown North\"],[\"4\",\"Midtown North\"],[\"5\",\"Midtown North\"],[\"6\",\"Midtown North\"],[\"7\",\"Midtown North\"],[\"8\",\"Midtown North\"],[\"9\",\"Midtown North\"],[\"0\",\"Midtown South\"],[\"1\",\"Midtown South\"],[\"2\",\"Midtown South\"],[\"3\",\"Midtown South\"],[\"4\",\"Midtown South\"],[\"5\",\"Midtown South\"],[\"6\",\"Midtown South\"],[\"7\",\"Midtown South\"],[\"8\",\"Midtown South\"],[\"9\",\"Midtown South\"],[\"0\",\"Midwood\"],[\"1\",\"Midwood\"],[\"2\",\"Midwood\"],[\"3\",\"Midwood\"],[\"4\",\"Midwood\"],[\"5\",\"Midwood\"],[\"6\",\"Midwood\"],[\"7\",\"Midwood\"],[\"8\",\"Midwood\"],[\"9\",\"Midwood\"],[\"0\",\"Morningside Heights\"],[\"1\",\"Morningside Heights\"],[\"2\",\"Morningside Heights\"],[\"3\",\"Morningside Heights\"],[\"4\",\"Morningside Heights\"],[\"5\",\"Morningside Heights\"],[\"6\",\"Morningside Heights\"],[\"7\",\"Morningside Heights\"],[\"8\",\"Morningside Heights\"],[\"9\",\"Morningside Heights\"],[\"0\",\"Morrisania/Melrose\"],[\"1\",\"Morrisania/Melrose\"],[\"2\",\"Morrisania/Melrose\"],[\"3\",\"Morrisania/Melrose\"],[\"4\",\"Morrisania/Melrose\"],[\"5\",\"Morrisania/Melrose\"],[\"6\",\"Morrisania/Melrose\"],[\"7\",\"Morrisania/Melrose\"],[\"8\",\"Morrisania/Melrose\"],[\"9\",\"Morrisania/Melrose\"],[\"0\",\"Mott Haven/Port Morris\"],[\"1\",\"Mott Haven/Port Morris\"],[\"2\",\"Mott Haven/Port Morris\"],[\"3\",\"Mott Haven/Port Morris\"],[\"4\",\"Mott Haven/Port Morris\"],[\"5\",\"Mott Haven/Port Morris\"],[\"6\",\"Mott Haven/Port Morris\"],[\"7\",\"Mott Haven/Port Morris\"],[\"8\",\"Mott Haven/Port Morris\"],[\"9\",\"Mott Haven/Port Morris\"],[\"0\",\"Mount Hope\"],[\"1\",\"Mount Hope\"],[\"2\",\"Mount Hope\"],[\"3\",\"Mount Hope\"],[\"4\",\"Mount Hope\"],[\"5\",\"Mount Hope\"],[\"6\",\"Mount Hope\"],[\"7\",\"Mount Hope\"],[\"8\",\"Mount Hope\"],[\"9\",\"Mount Hope\"],[\"0\",\"Murray Hill\"],[\"1\",\"Murray Hill\"],[\"2\",\"Murray Hill\"],[\"3\",\"Murray Hill\"],[\"4\",\"Murray Hill\"],[\"5\",\"Murray Hill\"],[\"6\",\"Murray Hill\"],[\"7\",\"Murray Hill\"],[\"8\",\"Murray Hill\"],[\"9\",\"Murray Hill\"],[\"0\",\"Murray Hill-Queens\"],[\"1\",\"Murray Hill-Queens\"],[\"2\",\"Murray Hill-Queens\"],[\"3\",\"Murray Hill-Queens\"],[\"4\",\"Murray Hill-Queens\"],[\"5\",\"Murray Hill-Queens\"],[\"6\",\"Murray Hill-Queens\"],[\"7\",\"Murray Hill-Queens\"],[\"8\",\"Murray Hill-Queens\"],[\"9\",\"Murray Hill-Queens\"],[\"0\",\"NA\"],[\"1\",\"NA\"],[\"2\",\"NA\"],[\"3\",\"NA\"],[\"4\",\"NA\"],[\"5\",\"NA\"],[\"6\",\"NA\"],[\"7\",\"NA\"],[\"8\",\"NA\"],[\"9\",\"NA\"],[\"0\",\"NV\"],[\"1\",\"NV\"],[\"2\",\"NV\"],[\"3\",\"NV\"],[\"4\",\"NV\"],[\"5\",\"NV\"],[\"6\",\"NV\"],[\"7\",\"NV\"],[\"8\",\"NV\"],[\"9\",\"NV\"],[\"0\",\"New Dorp/Midland Beach\"],[\"1\",\"New Dorp/Midland Beach\"],[\"2\",\"New Dorp/Midland Beach\"],[\"3\",\"New Dorp/Midland Beach\"],[\"4\",\"New Dorp/Midland Beach\"],[\"5\",\"New Dorp/Midland Beach\"],[\"6\",\"New Dorp/Midland Beach\"],[\"7\",\"New Dorp/Midland Beach\"],[\"8\",\"New Dorp/Midland Beach\"],[\"9\",\"New Dorp/Midland Beach\"],[\"0\",\"Newark Airport\"],[\"1\",\"Newark Airport\"],[\"2\",\"Newark Airport\"],[\"3\",\"Newark Airport\"],[\"4\",\"Newark Airport\"],[\"5\",\"Newark Airport\"],[\"6\",\"Newark Airport\"],[\"7\",\"Newark Airport\"],[\"8\",\"Newark Airport\"],[\"9\",\"Newark Airport\"],[\"0\",\"North Corona\"],[\"1\",\"North Corona\"],[\"2\",\"North Corona\"],[\"3\",\"North Corona\"],[\"4\",\"North Corona\"],[\"5\",\"North Corona\"],[\"6\",\"North Corona\"],[\"7\",\"North Corona\"],[\"8\",\"North Corona\"],[\"9\",\"North Corona\"],[\"0\",\"Norwood\"],[\"1\",\"Norwood\"],[\"2\",\"Norwood\"],[\"3\",\"Norwood\"],[\"4\",\"Norwood\"],[\"5\",\"Norwood\"],[\"6\",\"Norwood\"],[\"7\",\"Norwood\"],[\"8\",\"Norwood\"],[\"9\",\"Norwood\"],[\"0\",\"Oakland Gardens\"],[\"1\",\"Oakland Gardens\"],[\"2\",\"Oakland Gardens\"],[\"3\",\"Oakland Gardens\"],[\"4\",\"Oakland Gardens\"],[\"5\",\"Oakland Gardens\"],[\"6\",\"Oakland Gardens\"],[\"7\",\"Oakland Gardens\"],[\"8\",\"Oakland Gardens\"],[\"9\",\"Oakland Gardens\"],[\"0\",\"Oakwood\"],[\"1\",\"Oakwood\"],[\"2\",\"Oakwood\"],[\"3\",\"Oakwood\"],[\"4\",\"Oakwood\"],[\"5\",\"Oakwood\"],[\"6\",\"Oakwood\"],[\"7\",\"Oakwood\"],[\"8\",\"Oakwood\"],[\"9\",\"Oakwood\"],[\"0\",\"Ocean Hill\"],[\"1\",\"Ocean Hill\"],[\"2\",\"Ocean Hill\"],[\"3\",\"Ocean Hill\"],[\"4\",\"Ocean Hill\"],[\"5\",\"Ocean Hill\"],[\"6\",\"Ocean Hill\"],[\"7\",\"Ocean Hill\"],[\"8\",\"Ocean Hill\"],[\"9\",\"Ocean Hill\"],[\"0\",\"Ocean Parkway South\"],[\"1\",\"Ocean Parkway South\"],[\"2\",\"Ocean Parkway South\"],[\"3\",\"Ocean Parkway South\"],[\"4\",\"Ocean Parkway South\"],[\"5\",\"Ocean Parkway South\"],[\"6\",\"Ocean Parkway South\"],[\"7\",\"Ocean Parkway South\"],[\"8\",\"Ocean Parkway South\"],[\"9\",\"Ocean Parkway South\"],[\"0\",\"Old Astoria\"],[\"1\",\"Old Astoria\"],[\"2\",\"Old Astoria\"],[\"3\",\"Old Astoria\"],[\"4\",\"Old Astoria\"],[\"5\",\"Old Astoria\"],[\"6\",\"Old Astoria\"],[\"7\",\"Old Astoria\"],[\"8\",\"Old Astoria\"],[\"9\",\"Old Astoria\"],[\"0\",\"Ozone Park\"],[\"1\",\"Ozone Park\"],[\"2\",\"Ozone Park\"],[\"3\",\"Ozone Park\"],[\"4\",\"Ozone Park\"],[\"5\",\"Ozone Park\"],[\"6\",\"Ozone Park\"],[\"7\",\"Ozone Park\"],[\"8\",\"Ozone Park\"],[\"9\",\"Ozone Park\"],[\"0\",\"Park Slope\"],[\"1\",\"Park Slope\"],[\"2\",\"Park Slope\"],[\"3\",\"Park Slope\"],[\"4\",\"Park Slope\"],[\"5\",\"Park Slope\"],[\"6\",\"Park Slope\"],[\"7\",\"Park Slope\"],[\"8\",\"Park Slope\"],[\"9\",\"Park Slope\"],[\"0\",\"Parkchester\"],[\"1\",\"Parkchester\"],[\"2\",\"Parkchester\"],[\"3\",\"Parkchester\"],[\"4\",\"Parkchester\"],[\"5\",\"Parkchester\"],[\"6\",\"Parkchester\"],[\"7\",\"Parkchester\"],[\"8\",\"Parkchester\"],[\"9\",\"Parkchester\"],[\"0\",\"Pelham Bay\"],[\"1\",\"Pelham Bay\"],[\"2\",\"Pelham Bay\"],[\"3\",\"Pelham Bay\"],[\"4\",\"Pelham Bay\"],[\"5\",\"Pelham Bay\"],[\"6\",\"Pelham Bay\"],[\"7\",\"Pelham Bay\"],[\"8\",\"Pelham Bay\"],[\"9\",\"Pelham Bay\"],[\"0\",\"Pelham Bay Park\"],[\"1\",\"Pelham Bay Park\"],[\"2\",\"Pelham Bay Park\"],[\"3\",\"Pelham Bay Park\"],[\"4\",\"Pelham Bay Park\"],[\"5\",\"Pelham Bay Park\"],[\"6\",\"Pelham Bay Park\"],[\"7\",\"Pelham Bay Park\"],[\"8\",\"Pelham Bay Park\"],[\"9\",\"Pelham Bay Park\"],[\"0\",\"Pelham Parkway\"],[\"1\",\"Pelham Parkway\"],[\"2\",\"Pelham Parkway\"],[\"3\",\"Pelham Parkway\"],[\"4\",\"Pelham Parkway\"],[\"5\",\"Pelham Parkway\"],[\"6\",\"Pelham Parkway\"],[\"7\",\"Pelham Parkway\"],[\"8\",\"Pelham Parkway\"],[\"9\",\"Pelham Parkway\"],[\"0\",\"Penn Station/Madison Sq West\"],[\"1\",\"Penn Station/Madison Sq West\"],[\"2\",\"Penn Station/Madison Sq West\"],[\"3\",\"Penn Station/Madison Sq West\"],[\"4\",\"Penn Station/Madison Sq West\"],[\"5\",\"Penn Station/Madison Sq West\"],[\"6\",\"Penn Station/Madison Sq West\"],[\"7\",\"Penn Station/Madison Sq West\"],[\"8\",\"Penn Station/Madison Sq West\"],[\"9\",\"Penn Station/Madison Sq West\"],[\"0\",\"Port Richmond\"],[\"1\",\"Port Richmond\"],[\"2\",\"Port Richmond\"],[\"3\",\"Port Richmond\"],[\"4\",\"Port Richmond\"],[\"5\",\"Port Richmond\"],[\"6\",\"Port Richmond\"],[\"7\",\"Port Richmond\"],[\"8\",\"Port Richmond\"],[\"9\",\"Port Richmond\"],[\"0\",\"Prospect Heights\"],[\"1\",\"Prospect Heights\"],[\"2\",\"Prospect Heights\"],[\"3\",\"Prospect Heights\"],[\"4\",\"Prospect Heights\"],[\"5\",\"Prospect Heights\"],[\"6\",\"Prospect Heights\"],[\"7\",\"Prospect Heights\"],[\"8\",\"Prospect Heights\"],[\"9\",\"Prospect Heights\"],[\"0\",\"Prospect Park\"],[\"1\",\"Prospect Park\"],[\"2\",\"Prospect Park\"],[\"3\",\"Prospect Park\"],[\"4\",\"Prospect Park\"],[\"5\",\"Prospect Park\"],[\"6\",\"Prospect Park\"],[\"7\",\"Prospect Park\"],[\"8\",\"Prospect Park\"],[\"9\",\"Prospect Park\"],[\"0\",\"Prospect-Lefferts Gardens\"],[\"1\",\"Prospect-Lefferts Gardens\"],[\"2\",\"Prospect-Lefferts Gardens\"],[\"3\",\"Prospect-Lefferts Gardens\"],[\"4\",\"Prospect-Lefferts Gardens\"],[\"5\",\"Prospect-Lefferts Gardens\"],[\"6\",\"Prospect-Lefferts Gardens\"],[\"7\",\"Prospect-Lefferts Gardens\"],[\"8\",\"Prospect-Lefferts Gardens\"],[\"9\",\"Prospect-Lefferts Gardens\"],[\"0\",\"Queens Village\"],[\"1\",\"Queens Village\"],[\"2\",\"Queens Village\"],[\"3\",\"Queens Village\"],[\"4\",\"Queens Village\"],[\"5\",\"Queens Village\"],[\"6\",\"Queens Village\"],[\"7\",\"Queens Village\"],[\"8\",\"Queens Village\"],[\"9\",\"Queens Village\"],[\"0\",\"Queensboro Hill\"],[\"1\",\"Queensboro Hill\"],[\"2\",\"Queensboro Hill\"],[\"3\",\"Queensboro Hill\"],[\"4\",\"Queensboro Hill\"],[\"5\",\"Queensboro Hill\"],[\"6\",\"Queensboro Hill\"],[\"7\",\"Queensboro Hill\"],[\"8\",\"Queensboro Hill\"],[\"9\",\"Queensboro Hill\"],[\"0\",\"Queensbridge/Ravenswood\"],[\"1\",\"Queensbridge/Ravenswood\"],[\"2\",\"Queensbridge/Ravenswood\"],[\"3\",\"Queensbridge/Ravenswood\"],[\"4\",\"Queensbridge/Ravenswood\"],[\"5\",\"Queensbridge/Ravenswood\"],[\"6\",\"Queensbridge/Ravenswood\"],[\"7\",\"Queensbridge/Ravenswood\"],[\"8\",\"Queensbridge/Ravenswood\"],[\"9\",\"Queensbridge/Ravenswood\"],[\"0\",\"Randalls Island\"],[\"1\",\"Randalls Island\"],[\"2\",\"Randalls Island\"],[\"3\",\"Randalls Island\"],[\"4\",\"Randalls Island\"],[\"5\",\"Randalls Island\"],[\"6\",\"Randalls Island\"],[\"7\",\"Randalls Island\"],[\"8\",\"Randalls Island\"],[\"9\",\"Randalls Island\"],[\"0\",\"Red Hook\"],[\"1\",\"Red Hook\"],[\"2\",\"Red Hook\"],[\"3\",\"Red Hook\"],[\"4\",\"Red Hook\"],[\"5\",\"Red Hook\"],[\"6\",\"Red Hook\"],[\"7\",\"Red Hook\"],[\"8\",\"Red Hook\"],[\"9\",\"Red Hook\"],[\"0\",\"Rego Park\"],[\"1\",\"Rego Park\"],[\"2\",\"Rego Park\"],[\"3\",\"Rego Park\"],[\"4\",\"Rego Park\"],[\"5\",\"Rego Park\"],[\"6\",\"Rego Park\"],[\"7\",\"Rego Park\"],[\"8\",\"Rego Park\"],[\"9\",\"Rego Park\"],[\"0\",\"Richmond Hill\"],[\"1\",\"Richmond Hill\"],[\"2\",\"Richmond Hill\"],[\"3\",\"Richmond Hill\"],[\"4\",\"Richmond Hill\"],[\"5\",\"Richmond Hill\"],[\"6\",\"Richmond Hill\"],[\"7\",\"Richmond Hill\"],[\"8\",\"Richmond Hill\"],[\"9\",\"Richmond Hill\"],[\"0\",\"Ridgewood\"],[\"1\",\"Ridgewood\"],[\"2\",\"Ridgewood\"],[\"3\",\"Ridgewood\"],[\"4\",\"Ridgewood\"],[\"5\",\"Ridgewood\"],[\"6\",\"Ridgewood\"],[\"7\",\"Ridgewood\"],[\"8\",\"Ridgewood\"],[\"9\",\"Ridgewood\"],[\"0\",\"Rikers Island\"],[\"1\",\"Rikers Island\"],[\"2\",\"Rikers Island\"],[\"3\",\"Rikers Island\"],[\"4\",\"Rikers Island\"],[\"5\",\"Rikers Island\"],[\"6\",\"Rikers Island\"],[\"7\",\"Rikers Island\"],[\"8\",\"Rikers Island\"],[\"9\",\"Rikers Island\"],[\"0\",\"Riverdale/North Riverdale/Fieldston\"],[\"1\",\"Riverdale/North Riverdale/Fieldston\"],[\"2\",\"Riverdale/North Riverdale/Fieldston\"],[\"3\",\"Riverdale/North Riverdale/Fieldston\"],[\"4\",\"Riverdale/North Riverdale/Fieldston\"],[\"5\",\"Riverdale/North Riverdale/Fieldston\"],[\"6\",\"Riverdale/North Riverdale/Fieldston\"],[\"7\",\"Riverdale/North Riverdale/Fieldston\"],[\"8\",\"Riverdale/North Riverdale/Fieldston\"],[\"9\",\"Riverdale/North Riverdale/Fieldston\"],[\"0\",\"Rockaway Park\"],[\"1\",\"Rockaway Park\"],[\"2\",\"Rockaway Park\"],[\"3\",\"Rockaway Park\"],[\"4\",\"Rockaway Park\"],[\"5\",\"Rockaway Park\"],[\"6\",\"Rockaway Park\"],[\"7\",\"Rockaway Park\"],[\"8\",\"Rockaway Park\"],[\"9\",\"Rockaway Park\"],[\"0\",\"Roosevelt Island\"],[\"1\",\"Roosevelt Island\"],[\"2\",\"Roosevelt Island\"],[\"3\",\"Roosevelt Island\"],[\"4\",\"Roosevelt Island\"],[\"5\",\"Roosevelt Island\"],[\"6\",\"Roosevelt Island\"],[\"7\",\"Roosevelt Island\"],[\"8\",\"Roosevelt Island\"],[\"9\",\"Roosevelt Island\"],[\"0\",\"Rosedale\"],[\"1\",\"Rosedale\"],[\"2\",\"Rosedale\"],[\"3\",\"Rosedale\"],[\"4\",\"Rosedale\"],[\"5\",\"Rosedale\"],[\"6\",\"Rosedale\"],[\"7\",\"Rosedale\"],[\"8\",\"Rosedale\"],[\"9\",\"Rosedale\"],[\"0\",\"Rossville/Woodrow\"],[\"1\",\"Rossville/Woodrow\"],[\"2\",\"Rossville/Woodrow\"],[\"3\",\"Rossville/Woodrow\"],[\"4\",\"Rossville/Woodrow\"],[\"5\",\"Rossville/Woodrow\"],[\"6\",\"Rossville/Woodrow\"],[\"7\",\"Rossville/Woodrow\"],[\"8\",\"Rossville/Woodrow\"],[\"9\",\"Rossville/Woodrow\"],[\"0\",\"Saint Albans\"],[\"1\",\"Saint Albans\"],[\"2\",\"Saint Albans\"],[\"3\",\"Saint Albans\"],[\"4\",\"Saint Albans\"],[\"5\",\"Saint Albans\"],[\"6\",\"Saint Albans\"],[\"7\",\"Saint Albans\"],[\"8\",\"Saint Albans\"],[\"9\",\"Saint Albans\"],[\"0\",\"Saint George/New Brighton\"],[\"1\",\"Saint George/New Brighton\"],[\"2\",\"Saint George/New Brighton\"],[\"3\",\"Saint George/New Brighton\"],[\"4\",\"Saint George/New Brighton\"],[\"5\",\"Saint George/New Brighton\"],[\"6\",\"Saint George/New Brighton\"],[\"7\",\"Saint George/New Brighton\"],[\"8\",\"Saint George/New Brighton\"],[\"9\",\"Saint George/New Brighton\"],[\"0\",\"Saint Michaels Cemetery/Woodside\"],[\"1\",\"Saint Michaels Cemetery/Woodside\"],[\"2\",\"Saint Michaels Cemetery/Woodside\"],[\"3\",\"Saint Michaels Cemetery/Woodside\"],[\"4\",\"Saint Michaels Cemetery/Woodside\"],[\"5\",\"Saint Michaels Cemetery/Woodside\"],[\"6\",\"Saint Michaels Cemetery/Woodside\"],[\"7\",\"Saint Michaels Cemetery/Woodside\"],[\"8\",\"Saint Michaels Cemetery/Woodside\"],[\"9\",\"Saint Michaels Cemetery/Woodside\"],[\"0\",\"Schuylerville/Edgewater Park\"],[\"1\",\"Schuylerville/Edgewater Park\"],[\"2\",\"Schuylerville/Edgewater Park\"],[\"3\",\"Schuylerville/Edgewater Park\"],[\"4\",\"Schuylerville/Edgewater Park\"],[\"5\",\"Schuylerville/Edgewater Park\"],[\"6\",\"Schuylerville/Edgewater Park\"],[\"7\",\"Schuylerville/Edgewater Park\"],[\"8\",\"Schuylerville/Edgewater Park\"],[\"9\",\"Schuylerville/Edgewater Park\"],[\"0\",\"Seaport\"],[\"1\",\"Seaport\"],[\"2\",\"Seaport\"],[\"3\",\"Seaport\"],[\"4\",\"Seaport\"],[\"5\",\"Seaport\"],[\"6\",\"Seaport\"],[\"7\",\"Seaport\"],[\"8\",\"Seaport\"],[\"9\",\"Seaport\"],[\"0\",\"Sheepshead Bay\"],[\"1\",\"Sheepshead Bay\"],[\"2\",\"Sheepshead Bay\"],[\"3\",\"Sheepshead Bay\"],[\"4\",\"Sheepshead Bay\"],[\"5\",\"Sheepshead Bay\"],[\"6\",\"Sheepshead Bay\"],[\"7\",\"Sheepshead Bay\"],[\"8\",\"Sheepshead Bay\"],[\"9\",\"Sheepshead Bay\"],[\"0\",\"SoHo\"],[\"1\",\"SoHo\"],[\"2\",\"SoHo\"],[\"3\",\"SoHo\"],[\"4\",\"SoHo\"],[\"5\",\"SoHo\"],[\"6\",\"SoHo\"],[\"7\",\"SoHo\"],[\"8\",\"SoHo\"],[\"9\",\"SoHo\"],[\"0\",\"Soundview/Bruckner\"],[\"1\",\"Soundview/Bruckner\"],[\"2\",\"Soundview/Bruckner\"],[\"3\",\"Soundview/Bruckner\"],[\"4\",\"Soundview/Bruckner\"],[\"5\",\"Soundview/Bruckner\"],[\"6\",\"Soundview/Bruckner\"],[\"7\",\"Soundview/Bruckner\"],[\"8\",\"Soundview/Bruckner\"],[\"9\",\"Soundview/Bruckner\"],[\"0\",\"Soundview/Castle Hill\"],[\"1\",\"Soundview/Castle Hill\"],[\"2\",\"Soundview/Castle Hill\"],[\"3\",\"Soundview/Castle Hill\"],[\"4\",\"Soundview/Castle Hill\"],[\"5\",\"Soundview/Castle Hill\"],[\"6\",\"Soundview/Castle Hill\"],[\"7\",\"Soundview/Castle Hill\"],[\"8\",\"Soundview/Castle Hill\"],[\"9\",\"Soundview/Castle Hill\"],[\"0\",\"South Beach/Dongan Hills\"],[\"1\",\"South Beach/Dongan Hills\"],[\"2\",\"South Beach/Dongan Hills\"],[\"3\",\"South Beach/Dongan Hills\"],[\"4\",\"South Beach/Dongan Hills\"],[\"5\",\"South Beach/Dongan Hills\"],[\"6\",\"South Beach/Dongan Hills\"],[\"7\",\"South Beach/Dongan Hills\"],[\"8\",\"South Beach/Dongan Hills\"],[\"9\",\"South Beach/Dongan Hills\"],[\"0\",\"South Jamaica\"],[\"1\",\"South Jamaica\"],[\"2\",\"South Jamaica\"],[\"3\",\"South Jamaica\"],[\"4\",\"South Jamaica\"],[\"5\",\"South Jamaica\"],[\"6\",\"South Jamaica\"],[\"7\",\"South Jamaica\"],[\"8\",\"South Jamaica\"],[\"9\",\"South Jamaica\"],[\"0\",\"South Ozone Park\"],[\"1\",\"South Ozone Park\"],[\"2\",\"South Ozone Park\"],[\"3\",\"South Ozone Park\"],[\"4\",\"South Ozone Park\"],[\"5\",\"South Ozone Park\"],[\"6\",\"South Ozone Park\"],[\"7\",\"South Ozone Park\"],[\"8\",\"South Ozone Park\"],[\"9\",\"South Ozone Park\"],[\"0\",\"South Williamsburg\"],[\"1\",\"South Williamsburg\"],[\"2\",\"South Williamsburg\"],[\"3\",\"South Williamsburg\"],[\"4\",\"South Williamsburg\"],[\"5\",\"South Williamsburg\"],[\"6\",\"South Williamsburg\"],[\"7\",\"South Williamsburg\"],[\"8\",\"South Williamsburg\"],[\"9\",\"South Williamsburg\"],[\"0\",\"Springfield Gardens North\"],[\"1\",\"Springfield Gardens North\"],[\"2\",\"Springfield Gardens North\"],[\"3\",\"Springfield Gardens North\"],[\"4\",\"Springfield Gardens North\"],[\"5\",\"Springfield Gardens North\"],[\"6\",\"Springfield Gardens North\"],[\"7\",\"Springfield Gardens North\"],[\"8\",\"Springfield Gardens North\"],[\"9\",\"Springfield Gardens North\"],[\"0\",\"Springfield Gardens South\"],[\"1\",\"Springfield Gardens South\"],[\"2\",\"Springfield Gardens South\"],[\"3\",\"Springfield Gardens South\"],[\"4\",\"Springfield Gardens South\"],[\"5\",\"Springfield Gardens South\"],[\"6\",\"Springfield Gardens South\"],[\"7\",\"Springfield Gardens South\"],[\"8\",\"Springfield Gardens South\"],[\"9\",\"Springfield Gardens South\"],[\"0\",\"Spuyten Duyvil/Kingsbridge\"],[\"1\",\"Spuyten Duyvil/Kingsbridge\"],[\"2\",\"Spuyten Duyvil/Kingsbridge\"],[\"3\",\"Spuyten Duyvil/Kingsbridge\"],[\"4\",\"Spuyten Duyvil/Kingsbridge\"],[\"5\",\"Spuyten Duyvil/Kingsbridge\"],[\"6\",\"Spuyten Duyvil/Kingsbridge\"],[\"7\",\"Spuyten Duyvil/Kingsbridge\"],[\"8\",\"Spuyten Duyvil/Kingsbridge\"],[\"9\",\"Spuyten Duyvil/Kingsbridge\"],[\"0\",\"Stapleton\"],[\"1\",\"Stapleton\"],[\"2\",\"Stapleton\"],[\"3\",\"Stapleton\"],[\"4\",\"Stapleton\"],[\"5\",\"Stapleton\"],[\"6\",\"Stapleton\"],[\"7\",\"Stapleton\"],[\"8\",\"Stapleton\"],[\"9\",\"Stapleton\"],[\"0\",\"Starrett City\"],[\"1\",\"Starrett City\"],[\"2\",\"Starrett City\"],[\"3\",\"Starrett City\"],[\"4\",\"Starrett City\"],[\"5\",\"Starrett City\"],[\"6\",\"Starrett City\"],[\"7\",\"Starrett City\"],[\"8\",\"Starrett City\"],[\"9\",\"Starrett City\"],[\"0\",\"Steinway\"],[\"1\",\"Steinway\"],[\"2\",\"Steinway\"],[\"3\",\"Steinway\"],[\"4\",\"Steinway\"],[\"5\",\"Steinway\"],[\"6\",\"Steinway\"],[\"7\",\"Steinway\"],[\"8\",\"Steinway\"],[\"9\",\"Steinway\"],[\"0\",\"Stuy Town/Peter Cooper Village\"],[\"1\",\"Stuy Town/Peter Cooper Village\"],[\"2\",\"Stuy Town/Peter Cooper Village\"],[\"3\",\"Stuy Town/Peter Cooper Village\"],[\"4\",\"Stuy Town/Peter Cooper Village\"],[\"5\",\"Stuy Town/Peter Cooper Village\"],[\"6\",\"Stuy Town/Peter Cooper Village\"],[\"7\",\"Stuy Town/Peter Cooper Village\"],[\"8\",\"Stuy Town/Peter Cooper Village\"],[\"9\",\"Stuy Town/Peter Cooper Village\"],[\"0\",\"Stuyvesant Heights\"],[\"1\",\"Stuyvesant Heights\"],[\"2\",\"Stuyvesant Heights\"],[\"3\",\"Stuyvesant Heights\"],[\"4\",\"Stuyvesant Heights\"],[\"5\",\"Stuyvesant Heights\"],[\"6\",\"Stuyvesant Heights\"],[\"7\",\"Stuyvesant Heights\"],[\"8\",\"Stuyvesant Heights\"],[\"9\",\"Stuyvesant Heights\"],[\"0\",\"Sunnyside\"],[\"1\",\"Sunnyside\"],[\"2\",\"Sunnyside\"],[\"3\",\"Sunnyside\"],[\"4\",\"Sunnyside\"],[\"5\",\"Sunnyside\"],[\"6\",\"Sunnyside\"],[\"7\",\"Sunnyside\"],[\"8\",\"Sunnyside\"],[\"9\",\"Sunnyside\"],[\"0\",\"Sunset Park East\"],[\"1\",\"Sunset Park East\"],[\"2\",\"Sunset Park East\"],[\"3\",\"Sunset Park East\"],[\"4\",\"Sunset Park East\"],[\"5\",\"Sunset Park East\"],[\"6\",\"Sunset Park East\"],[\"7\",\"Sunset Park East\"],[\"8\",\"Sunset Park East\"],[\"9\",\"Sunset Park East\"],[\"0\",\"Sunset Park West\"],[\"1\",\"Sunset Park West\"],[\"2\",\"Sunset Park West\"],[\"3\",\"Sunset Park West\"],[\"4\",\"Sunset Park West\"],[\"5\",\"Sunset Park West\"],[\"6\",\"Sunset Park West\"],[\"7\",\"Sunset Park West\"],[\"8\",\"Sunset Park West\"],[\"9\",\"Sunset Park West\"],[\"0\",\"Sutton Place/Turtle Bay North\"],[\"1\",\"Sutton Place/Turtle Bay North\"],[\"2\",\"Sutton Place/Turtle Bay North\"],[\"3\",\"Sutton Place/Turtle Bay North\"],[\"4\",\"Sutton Place/Turtle Bay North\"],[\"5\",\"Sutton Place/Turtle Bay North\"],[\"6\",\"Sutton Place/Turtle Bay North\"],[\"7\",\"Sutton Place/Turtle Bay North\"],[\"8\",\"Sutton Place/Turtle Bay North\"],[\"9\",\"Sutton Place/Turtle Bay North\"],[\"0\",\"Times Sq/Theatre District\"],[\"1\",\"Times Sq/Theatre District\"],[\"2\",\"Times Sq/Theatre District\"],[\"3\",\"Times Sq/Theatre District\"],[\"4\",\"Times Sq/Theatre District\"],[\"5\",\"Times Sq/Theatre District\"],[\"6\",\"Times Sq/Theatre District\"],[\"7\",\"Times Sq/Theatre District\"],[\"8\",\"Times Sq/Theatre District\"],[\"9\",\"Times Sq/Theatre District\"],[\"0\",\"TriBeCa/Civic Center\"],[\"1\",\"TriBeCa/Civic Center\"],[\"2\",\"TriBeCa/Civic Center\"],[\"3\",\"TriBeCa/Civic Center\"],[\"4\",\"TriBeCa/Civic Center\"],[\"5\",\"TriBeCa/Civic Center\"],[\"6\",\"TriBeCa/Civic Center\"],[\"7\",\"TriBeCa/Civic Center\"],[\"8\",\"TriBeCa/Civic Center\"],[\"9\",\"TriBeCa/Civic Center\"],[\"0\",\"Two Bridges/Seward Park\"],[\"1\",\"Two Bridges/Seward Park\"],[\"2\",\"Two Bridges/Seward Park\"],[\"3\",\"Two Bridges/Seward Park\"],[\"4\",\"Two Bridges/Seward Park\"],[\"5\",\"Two Bridges/Seward Park\"],[\"6\",\"Two Bridges/Seward Park\"],[\"7\",\"Two Bridges/Seward Park\"],[\"8\",\"Two Bridges/Seward Park\"],[\"9\",\"Two Bridges/Seward Park\"],[\"0\",\"UN/Turtle Bay South\"],[\"1\",\"UN/Turtle Bay South\"],[\"2\",\"UN/Turtle Bay South\"],[\"3\",\"UN/Turtle Bay South\"],[\"4\",\"UN/Turtle Bay South\"],[\"5\",\"UN/Turtle Bay South\"],[\"6\",\"UN/Turtle Bay South\"],[\"7\",\"UN/Turtle Bay South\"],[\"8\",\"UN/Turtle Bay South\"],[\"9\",\"UN/Turtle Bay South\"],[\"0\",\"Union Sq\"],[\"1\",\"Union Sq\"],[\"2\",\"Union Sq\"],[\"3\",\"Union Sq\"],[\"4\",\"Union Sq\"],[\"5\",\"Union Sq\"],[\"6\",\"Union Sq\"],[\"7\",\"Union Sq\"],[\"8\",\"Union Sq\"],[\"9\",\"Union Sq\"],[\"0\",\"University Heights/Morris Heights\"],[\"1\",\"University Heights/Morris Heights\"],[\"2\",\"University Heights/Morris Heights\"],[\"3\",\"University Heights/Morris Heights\"],[\"4\",\"University Heights/Morris Heights\"],[\"5\",\"University Heights/Morris Heights\"],[\"6\",\"University Heights/Morris Heights\"],[\"7\",\"University Heights/Morris Heights\"],[\"8\",\"University Heights/Morris Heights\"],[\"9\",\"University Heights/Morris Heights\"],[\"0\",\"Upper East Side North\"],[\"1\",\"Upper East Side North\"],[\"2\",\"Upper East Side North\"],[\"3\",\"Upper East Side North\"],[\"4\",\"Upper East Side North\"],[\"5\",\"Upper East Side North\"],[\"6\",\"Upper East Side North\"],[\"7\",\"Upper East Side North\"],[\"8\",\"Upper East Side North\"],[\"9\",\"Upper East Side North\"],[\"0\",\"Upper East Side South\"],[\"1\",\"Upper East Side South\"],[\"2\",\"Upper East Side South\"],[\"3\",\"Upper East Side South\"],[\"4\",\"Upper East Side South\"],[\"5\",\"Upper East Side South\"],[\"6\",\"Upper East Side South\"],[\"7\",\"Upper East Side South\"],[\"8\",\"Upper East Side South\"],[\"9\",\"Upper East Side South\"],[\"0\",\"Upper West Side North\"],[\"1\",\"Upper West Side North\"],[\"2\",\"Upper West Side North\"],[\"3\",\"Upper West Side North\"],[\"4\",\"Upper West Side North\"],[\"5\",\"Upper West Side North\"],[\"6\",\"Upper West Side North\"],[\"7\",\"Upper West Side North\"],[\"8\",\"Upper West Side North\"],[\"9\",\"Upper West Side North\"],[\"0\",\"Upper West Side South\"],[\"1\",\"Upper West Side South\"],[\"2\",\"Upper West Side South\"],[\"3\",\"Upper West Side South\"],[\"4\",\"Upper West Side South\"],[\"5\",\"Upper West Side South\"],[\"6\",\"Upper West Side South\"],[\"7\",\"Upper West Side South\"],[\"8\",\"Upper West Side South\"],[\"9\",\"Upper West Side South\"],[\"0\",\"Van Cortlandt Park\"],[\"1\",\"Van Cortlandt Park\"],[\"2\",\"Van Cortlandt Park\"],[\"3\",\"Van Cortlandt Park\"],[\"4\",\"Van Cortlandt Park\"],[\"5\",\"Van Cortlandt Park\"],[\"6\",\"Van Cortlandt Park\"],[\"7\",\"Van Cortlandt Park\"],[\"8\",\"Van Cortlandt Park\"],[\"9\",\"Van Cortlandt Park\"],[\"0\",\"Van Cortlandt Village\"],[\"1\",\"Van Cortlandt Village\"],[\"2\",\"Van Cortlandt Village\"],[\"3\",\"Van Cortlandt Village\"],[\"4\",\"Van Cortlandt Village\"],[\"5\",\"Van Cortlandt Village\"],[\"6\",\"Van Cortlandt Village\"],[\"7\",\"Van Cortlandt Village\"],[\"8\",\"Van Cortlandt Village\"],[\"9\",\"Van Cortlandt Village\"],[\"0\",\"Van Nest/Morris Park\"],[\"1\",\"Van Nest/Morris Park\"],[\"2\",\"Van Nest/Morris Park\"],[\"3\",\"Van Nest/Morris Park\"],[\"4\",\"Van Nest/Morris Park\"],[\"5\",\"Van Nest/Morris Park\"],[\"6\",\"Van Nest/Morris Park\"],[\"7\",\"Van Nest/Morris Park\"],[\"8\",\"Van Nest/Morris Park\"],[\"9\",\"Van Nest/Morris Park\"],[\"0\",\"Washington Heights North\"],[\"1\",\"Washington Heights North\"],[\"2\",\"Washington Heights North\"],[\"3\",\"Washington Heights North\"],[\"4\",\"Washington Heights North\"],[\"5\",\"Washington Heights North\"],[\"6\",\"Washington Heights North\"],[\"7\",\"Washington Heights North\"],[\"8\",\"Washington Heights North\"],[\"9\",\"Washington Heights North\"],[\"0\",\"Washington Heights South\"],[\"1\",\"Washington Heights South\"],[\"2\",\"Washington Heights South\"],[\"3\",\"Washington Heights South\"],[\"4\",\"Washington Heights South\"],[\"5\",\"Washington Heights South\"],[\"6\",\"Washington Heights South\"],[\"7\",\"Washington Heights South\"],[\"8\",\"Washington Heights South\"],[\"9\",\"Washington Heights South\"],[\"0\",\"West Brighton\"],[\"1\",\"West Brighton\"],[\"2\",\"West Brighton\"],[\"3\",\"West Brighton\"],[\"4\",\"West Brighton\"],[\"5\",\"West Brighton\"],[\"6\",\"West Brighton\"],[\"7\",\"West Brighton\"],[\"8\",\"West Brighton\"],[\"9\",\"West Brighton\"],[\"0\",\"West Chelsea/Hudson Yards\"],[\"1\",\"West Chelsea/Hudson Yards\"],[\"2\",\"West Chelsea/Hudson Yards\"],[\"3\",\"West Chelsea/Hudson Yards\"],[\"4\",\"West Chelsea/Hudson Yards\"],[\"5\",\"West Chelsea/Hudson Yards\"],[\"6\",\"West Chelsea/Hudson Yards\"],[\"7\",\"West Chelsea/Hudson Yards\"],[\"8\",\"West Chelsea/Hudson Yards\"],[\"9\",\"West Chelsea/Hudson Yards\"],[\"0\",\"West Concourse\"],[\"1\",\"West Concourse\"],[\"2\",\"West Concourse\"],[\"3\",\"West Concourse\"],[\"4\",\"West Concourse\"],[\"5\",\"West Concourse\"],[\"6\",\"West Concourse\"],[\"7\",\"West Concourse\"],[\"8\",\"West Concourse\"],[\"9\",\"West Concourse\"],[\"0\",\"West Farms/Bronx River\"],[\"1\",\"West Farms/Bronx River\"],[\"2\",\"West Farms/Bronx River\"],[\"3\",\"West Farms/Bronx River\"],[\"4\",\"West Farms/Bronx River\"],[\"5\",\"West Farms/Bronx River\"],[\"6\",\"West Farms/Bronx River\"],[\"7\",\"West Farms/Bronx River\"],[\"8\",\"West Farms/Bronx River\"],[\"9\",\"West Farms/Bronx River\"],[\"0\",\"West Village\"],[\"1\",\"West Village\"],[\"2\",\"West Village\"],[\"3\",\"West Village\"],[\"4\",\"West Village\"],[\"5\",\"West Village\"],[\"6\",\"West Village\"],[\"7\",\"West Village\"],[\"8\",\"West Village\"],[\"9\",\"West Village\"],[\"0\",\"Westchester Village/Unionport\"],[\"1\",\"Westchester Village/Unionport\"],[\"2\",\"Westchester Village/Unionport\"],[\"3\",\"Westchester Village/Unionport\"],[\"4\",\"Westchester Village/Unionport\"],[\"5\",\"Westchester Village/Unionport\"],[\"6\",\"Westchester Village/Unionport\"],[\"7\",\"Westchester Village/Unionport\"],[\"8\",\"Westchester Village/Unionport\"],[\"9\",\"Westchester Village/Unionport\"],[\"0\",\"Westerleigh\"],[\"1\",\"Westerleigh\"],[\"2\",\"Westerleigh\"],[\"3\",\"Westerleigh\"],[\"4\",\"Westerleigh\"],[\"5\",\"Westerleigh\"],[\"6\",\"Westerleigh\"],[\"7\",\"Westerleigh\"],[\"8\",\"Westerleigh\"],[\"9\",\"Westerleigh\"],[\"0\",\"Whitestone\"],[\"1\",\"Whitestone\"],[\"2\",\"Whitestone\"],[\"3\",\"Whitestone\"],[\"4\",\"Whitestone\"],[\"5\",\"Whitestone\"],[\"6\",\"Whitestone\"],[\"7\",\"Whitestone\"],[\"8\",\"Whitestone\"],[\"9\",\"Whitestone\"],[\"0\",\"Willets Point\"],[\"1\",\"Willets Point\"],[\"2\",\"Willets Point\"],[\"3\",\"Willets Point\"],[\"4\",\"Willets Point\"],[\"5\",\"Willets Point\"],[\"6\",\"Willets Point\"],[\"7\",\"Willets Point\"],[\"8\",\"Willets Point\"],[\"9\",\"Willets Point\"],[\"0\",\"Williamsbridge/Olinville\"],[\"1\",\"Williamsbridge/Olinville\"],[\"2\",\"Williamsbridge/Olinville\"],[\"3\",\"Williamsbridge/Olinville\"],[\"4\",\"Williamsbridge/Olinville\"],[\"5\",\"Williamsbridge/Olinville\"],[\"6\",\"Williamsbridge/Olinville\"],[\"7\",\"Williamsbridge/Olinville\"],[\"8\",\"Williamsbridge/Olinville\"],[\"9\",\"Williamsbridge/Olinville\"],[\"0\",\"Williamsburg (North Side)\"],[\"1\",\"Williamsburg (North Side)\"],[\"2\",\"Williamsburg (North Side)\"],[\"3\",\"Williamsburg (North Side)\"],[\"4\",\"Williamsburg (North Side)\"],[\"5\",\"Williamsburg (North Side)\"],[\"6\",\"Williamsburg (North Side)\"],[\"7\",\"Williamsburg (North Side)\"],[\"8\",\"Williamsburg (North Side)\"],[\"9\",\"Williamsburg (North Side)\"],[\"0\",\"Williamsburg (South Side)\"],[\"1\",\"Williamsburg (South Side)\"],[\"2\",\"Williamsburg (South Side)\"],[\"3\",\"Williamsburg (South Side)\"],[\"4\",\"Williamsburg (South Side)\"],[\"5\",\"Williamsburg (South Side)\"],[\"6\",\"Williamsburg (South Side)\"],[\"7\",\"Williamsburg (South Side)\"],[\"8\",\"Williamsburg (South Side)\"],[\"9\",\"Williamsburg (South Side)\"],[\"0\",\"Windsor Terrace\"],[\"1\",\"Windsor Terrace\"],[\"2\",\"Windsor Terrace\"],[\"3\",\"Windsor Terrace\"],[\"4\",\"Windsor Terrace\"],[\"5\",\"Windsor Terrace\"],[\"6\",\"Windsor Terrace\"],[\"7\",\"Windsor Terrace\"],[\"8\",\"Windsor Terrace\"],[\"9\",\"Windsor Terrace\"],[\"0\",\"Woodhaven\"],[\"1\",\"Woodhaven\"],[\"2\",\"Woodhaven\"],[\"3\",\"Woodhaven\"],[\"4\",\"Woodhaven\"],[\"5\",\"Woodhaven\"],[\"6\",\"Woodhaven\"],[\"7\",\"Woodhaven\"],[\"8\",\"Woodhaven\"],[\"9\",\"Woodhaven\"],[\"0\",\"Woodside\"],[\"1\",\"Woodside\"],[\"2\",\"Woodside\"],[\"3\",\"Woodside\"],[\"4\",\"Woodside\"],[\"5\",\"Woodside\"],[\"6\",\"Woodside\"],[\"7\",\"Woodside\"],[\"8\",\"Woodside\"],[\"9\",\"Woodside\"],[\"0\",\"World Trade Center\"],[\"1\",\"World Trade Center\"],[\"2\",\"World Trade Center\"],[\"3\",\"World Trade Center\"],[\"4\",\"World Trade Center\"],[\"5\",\"World Trade Center\"],[\"6\",\"World Trade Center\"],[\"7\",\"World Trade Center\"],[\"8\",\"World Trade Center\"],[\"9\",\"World Trade Center\"],[\"0\",\"Yorkville East\"],[\"1\",\"Yorkville East\"],[\"2\",\"Yorkville East\"],[\"3\",\"Yorkville East\"],[\"4\",\"Yorkville East\"],[\"5\",\"Yorkville East\"],[\"6\",\"Yorkville East\"],[\"7\",\"Yorkville East\"],[\"8\",\"Yorkville East\"],[\"9\",\"Yorkville East\"],[\"0\",\"Yorkville West\"],[\"1\",\"Yorkville West\"],[\"2\",\"Yorkville West\"],[\"3\",\"Yorkville West\"],[\"4\",\"Yorkville West\"],[\"5\",\"Yorkville West\"],[\"6\",\"Yorkville West\"],[\"7\",\"Yorkville West\"],[\"8\",\"Yorkville West\"],[\"9\",\"Yorkville West\"]]",
	//    "logName": "1554760860756237108-dG90YWwgZmFyZSBhbW91bnQgYnkgcGFzc2VuZ2VyIGNvdW50IGZyb20gem9uZSB0byBib3JvdWdoIGxhc3QgeWVhcg==",
	//    "questionSuccess": "true",
	//    "refDate": "",
	//    "rollup": "{\"token\":\"daily\"}",
	//    "totalGraphs": "2610",
	//    "ts": "{\"start\":\"2018-01-01T00:00:00Z\",\"end\":\"2018-12-31T00:00:00Z\"}"
	//  },
	//  "dialogAction": {
	//    "type": "Close",
	//    "fulfillmentState": "Fulfilled",
	//    "message": {
	//      "content": "We found you the following graph:",
	//      "contentType": "PlainText"
	//    }
	//  }
	//}`

	// Test with
	// base65536
	// base122

	f, err := ioutil.ReadFile("/Users/Peter/go/src/github.com/ptimson/lz-string-go/peter/data2")
	data2 := string(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data")
	fmt.Println(len(data2))
	fmt.Println("--------------------------")

	fmt.Println("LZ-String Base64")
	startTime := time.Now()
	//fmt.Println(CompressToBase64(data2))
	s := CompressToBase64(data2)
	fmt.Println(len(s))
	fmt.Println(time.Since(startTime))

	fmt.Println("--------------------------")
	fmt.Println("LZ-String UTF16")
	startTime = time.Now()
	//fmt.Println(CompressToUTF16(data2))
	s = CompressToUTF16(data2)
	fmt.Println(len(s))
	fmt.Println(time.Since(startTime))

	fmt.Println("--------------------------")
	fmt.Println("GZIP Base64")
	startTime = time.Now()
	b := &bytes.Buffer{}
	g := gzip.NewWriter(b)
	g.Write([]byte(data2))
	g.Flush()
	//fmt.Println(string(base64.StdEncoding.EncodeToString(b.Bytes())))
	fmt.Println(len(b.Bytes()))
	fmt.Println(time.Since(startTime))

	fmt.Println("--------------------------")
	fmt.Println("ZLIB Base64")
	startTime = time.Now()
	b = &bytes.Buffer{}
	z := zlib.NewWriter(b)
	z.Write([]byte(data2))
	z.Flush()
	//fmt.Println(string(base64.StdEncoding.EncodeToString(b.Bytes())))
	fmt.Println(len(b.Bytes()))
	fmt.Println(time.Since(startTime))

	fmt.Println("--------------------------")
	fmt.Println("LZMA Base64")
	startTime = time.Now()
	b = &bytes.Buffer{}
	l := lzma.NewWriterLevel(b, 1)
	l.Write([]byte(data2))
	l.Close()
	//fmt.Println(string(base64.StdEncoding.EncodeToString(b.Bytes())))
	fmt.Println(len(b.Bytes()))
	fmt.Println(time.Since(startTime))

}

const keyStrBase64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
const keyStrUriSafe = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-$"

var baseReverseDic = make(map[string]map[rune]int)

func getBaseValue(alphabet string, char rune) (int, bool) {
	_, ok := baseReverseDic[alphabet]
	if !ok {
		baseReverseDic[alphabet] = make(map[rune]int, len(alphabet))
		for i, c := range alphabet {
			baseReverseDic[alphabet][c] = i
		}
	}
	value, ok := baseReverseDic[alphabet][char]
	return value, ok
}

func CompressToBase64(input string) string {
	if input == "" {
		return ""
	}

	res := compress(input, 6, func(i int) rune {
		return rune(keyStrBase64[i])
	})

	switch len(res) % 4 {
	case 1:
		return res + "==="
	case 2:
		return res + "=="
	case 3:
		return res + "="
	}

	return res
}

//CompressToEncodedURIComponent compress into a string that is already URI encoded
func CompressToEncodedURIComponent(input string) string {
	if input == "" {
		return ""
	}
	return compress(input, 6, func(i int) rune {
		return rune(keyStrUriSafe[i])
	})
}

// CompressToUInt8Array compress into uint8array (UCS-2 big endian format)
func CompressToUInt8Array(input string) []uint8 {
	var compressed = Compress(input)
	buf := make([]uint8, len(compressed)*2)

	for i := 0; i < len(compressed); i++ {
		r := compressed[i]
		buf[i*2] = uint8(r) >> uint8(8)
		tmp := int(r) % 256
		buf[i*2+1] = uint8(tmp)
	}

	return buf
}

func CompressToUTF16(input string) string {
	if input == "" {
		return ""
	}
	return compress(input, 15, func(i int) rune {
		return rune(i + 32)
	}) + " "
}

func Compress(input string) string {
	return compress(input, 16, func(i int) rune {
		return rune(i)
	})
}

func decompressFromBase64(input string) string {
	if input == "" {
		return ""
	}

	return decompress(len(input), 32, func(i int) (int, bool) {
		return getBaseValue(keyStrBase64, rune(keyStrBase64[i]))
	})
}

func decompress(length int, A int, f func(i int) (int, bool)) string {
	return ""
}

type context struct {
	key                string
	val                int
	position           int
	enlargeIn          float64
	bits               int
	dictSize           int
	dictionary         map[string]int
	dictionaryToCreate map[string]struct{}
	data               []rune
}

func compress(uncompressed string, bitsPerChar int, getCharFromInt func(i int) rune) string {
	if uncompressed == "" {
		return ""
	}

	w := ""                 // key
	dictSize := 3           // dictSize
	numBits := 2            // bits
	dataVal := 0            // val
	dataPosition := 0       // position
	enlargeIn := float64(2) // Compensate for the first entry which should not count
	var data []rune

	dictionary := make(map[string]int)
	dictionaryToCreate := make(map[string]struct{})

	for _, c := range uncompressed {
		char := string(c)
		_, ok := dictionary[char]
		if !ok {
			dictionary[char] = dictSize
			dictSize++
			dictionaryToCreate[char] = struct{}{}
		}

		wc := fmt.Sprintf("%s%s", w, char)
		_, ok = dictionary[wc]
		if ok {
			w = wc
			continue
		}

		_, ok = dictionaryToCreate[w]
		if !ok {
			value, ok := dictionary[w]
			if !ok {
				value = 0
			}
			for i := 0; i < numBits; i++ {
				dataVal = (dataVal << 1) | (value & 1)
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = value >> 1
			}
		} else {

			var max255 uint8 = 255
			if len(w) > 0 && w[0] <= max255 {
				for i := 0; i < numBits; i++ {
					dataVal = dataVal << 1
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
				}
				value := int(w[0])
				for i := 0; i < 8; i++ {
					dataVal = dataVal<<1 | (value & 1)
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
					value = value >> 1
				}

			} else {

				value := 1
				for i := 0; i < numBits; i++ {
					dataVal = (dataVal << 1) | value
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
					value = 0
				}

				value = int(w[0])

				for i := 0; i < 16; i++ {
					dataVal = dataVal<<1 | (value & 1)
					if dataPosition == bitsPerChar-1 {
						data = append(data, getCharFromInt(dataVal))
						dataPosition = 0
						dataVal = 0
					} else {
						dataPosition++
					}
					value = value >> 1
				}
			}
			enlargeIn--
			if enlargeIn == 0 {
				enlargeIn = math.Pow(2, float64(numBits))
				numBits++
			}
			delete(dictionaryToCreate, w)
		}

		enlargeIn--
		if enlargeIn == 0 {
			enlargeIn = math.Pow(2, float64(numBits))
			numBits++
		}
		// Add wc to the dictionary.
		dictionary[wc] = dictSize
		dictSize++
		w = char
	}

	/////////
	// Output the code for w.
	/////////
	_, ok := dictionaryToCreate[w]
	if w != "" && ok {
		var max255 uint8 = 255
		if len(w) > 0 && w[0] <= max255 {
			for i := 0; i < numBits; i++ {
				dataVal = dataVal << 1
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
			}
			value := int(w[0])
			for i := 0; i < 8; i++ {
				dataVal = dataVal<<1 | (value & 1)
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = value >> 1
			}

		} else {

			value := 1
			for i := 0; i < numBits; i++ {
				dataVal = (dataVal << 1) | value
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = 0
			}

			value = int(w[0])

			for i := 0; i < 16; i++ {
				dataVal = dataVal<<1 | (value & 1)
				if dataPosition == bitsPerChar-1 {
					data = append(data, getCharFromInt(dataVal))
					dataPosition = 0
					dataVal = 0
				} else {
					dataPosition++
				}
				value = value >> 1
			}
		}

		enlargeIn--
		if enlargeIn == 0 {
			enlargeIn = math.Pow(2, float64(numBits))
			numBits++
		}
		delete(dictionaryToCreate, w)
	}

	_, ok = dictionaryToCreate[w]
	if w != "" && !ok {
		value, ok := dictionary[w]
		if !ok {
			value = 0
		}
		for i := 0; i < numBits; i++ {
			dataVal = dataVal<<1 | (value & 1)
			if dataPosition == bitsPerChar-1 {
				data = append(data, getCharFromInt(dataVal))
				dataPosition = 0
				dataVal = 0
			} else {
				dataPosition++
			}
			value = value >> 1
		}
	}

	if w != "" {
		enlargeIn--
		if enlargeIn == 0 {
			enlargeIn = math.Pow(2, float64(numBits))
			numBits++
		}
	}

	// Mark the end of the stream
	value := 2
	for i := 0; i < numBits; i++ {
		dataVal = (dataVal << 1) | (value & 1)
		if dataPosition == bitsPerChar-1 {
			data = append(data, getCharFromInt(dataVal))
			dataPosition = 0
			dataVal = 0
		} else {
			dataPosition++
		}
		value = value >> 1
	}

	// Flush the last char
	for {
		dataVal = dataVal << 1
		if dataPosition == bitsPerChar-1 {
			data = append(data, getCharFromInt(dataVal))
			break
		} else {
			dataPosition++
		}
	}

	return string(data)
}
