<script type="text/javascript" charset="utf-8">



        $(document).ready(function(){

            CountryCounter = {}
            Countries = {}
            RTMapNamespace = '/rtmap'; // change to an empty string to use the global namespace

            // the socket.io documentation recommends sending an explicit package upon connection
            // this is specially important when using the global namespace
            var rtsocket = io.connect('http://' + document.domain + ':' + location.port + RTMapNamespace, {"force new connection": false});

            // New sample received from
            rtsocket.on('update', function(data) {

            });

            // event handler for new connections
            rtsocket.on('connect', function(data) {
                if ( data !== undefined )
                {
                    parseRTData(JSON.parse(data));
                }
            });
        });

        function parseRTData(data) {
            $.each(data, function(index, item){
                $.each(item.ipMeta, function(idx, ip) {
                    if ( ip.iso_code in CountryCounter )
                    {
                        CountryCounter[ip.iso_code] += 1;

                    }
                    else
                    {
                        CountryCounter[ip.iso_code] = 1;
                        Countries[ip.iso_code] = ip.country;
                    }
                });
            });
            $.each(CountryCounter, function(country, count){
                mapData.push({
                    "code":country,
                    "name": Countries[country],
                    "value": count,
                    "color": getRandomColor()
                });
            });
            if ( map !== undefined)
            {
                map.dataProvider = createCircles(mapData, map.dataProvider);
                map.validateData;
            }
        };
</script>
