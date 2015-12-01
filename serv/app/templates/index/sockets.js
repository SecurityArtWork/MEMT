<script type="text/javascript" charset="utf-8">
        $(document).ready(function(){
            CountryCounter = {}
            Countries = {}
            RTMapNamespace = '/rtmap';

            // the socket.io documentation recommends sending an explicit package upon connection
            // this is specially important when using the global namespace
            var rtsocket = io.connect('http://' + document.domain + ':' + location.port + RTMapNamespace, {"force new connection": false});

            // New sample received from
            rtsocket.on('update', function(data) {
                console.debug("UPDATING...");
                parseRTData(JSON.parse(data));
            });

            // event handler for new connections
            rtsocket.on('connect', function(data) {
                console.debug("CONNECTED...");
                if ( data !== undefined )
                {
                    console.debug("CONNECTED... GOT DATA");
                    parseRTData(JSON.parse(data));
                } else{
                    console.debug("CONNECTED... NO DATA");
                }
            });
        });
        //Let's parse the JSON data for each item
        function parseRTData(data) {
            $.each(data, function(index, item){
                $.each(item.ipmeta, function(idx, ip) {
                    if ( ip.iso_code in CountryCounter )
                    {
                        //increment the number of samples uploaded from this country
                        CountryCounter[ip.iso_code] += 1;
                    }
                    else
                    {
                        //setting the counter for this country
                        CountryCounter[ip.iso_code] = 1;
                        //Assign the country name to the iso_code
                        Countries[ip.iso_code] = ip.country;
                    }
                });
            });
            //Pushing data to mapData array which is going to be represented on RT Map.
            $.each(CountryCounter, function(country, count){
                //checking if we are going to be able to represent the country
                if ( country !== "null" && country !== "unknown")
                {
                    color = getRandomColor();
                    mapData.push({
                        "code": country,
                        "name": Countries[country],
                        "value": count,
                        "color": color
                    });
                }
            });
            if ( map !== undefined)
            {
                //if map is already loaded, we only have to recalculate circles and validateData.
                map.dataProvider = createCircles(mapData, dataProvider);
                map.validateData();
            }
            //Set counter label to show the number of countries painted.
            jQuery("#countriesInMap").text("Showing "+mapData.length+" countries");

        };
</script>
