<script type="text/javascript" charset="utf-8">
        $(document).ready(function(){
            RTMapNamespace = '/feed'; // change to an empty string to use the global namespace

            // the socket.io documentation recommends sending an explicit package upon connection
            // this is specially important when using the global namespace
            var rtsocket = io.connect('http://' + document.domain + ':' + location.port + RTMapNamespace);

            // event handler for server sent data
            // the data is displayed in the "Received" section of the page
            socket.on('update', function(data) {
                parseData(data);
            });

            // event handler for new connections
            socket.on('connect', function(data) {
                parseData(data);
            });
        });

        function parseData(data){
            feed_ul = $('#feed');
            $.each(JSON.parse(data), function(idx, obj) {
                feed_li = feed_ul.append($('<li>').append(
                    $('<i>').attr('class', 'fa fa-fw fa-bullhorn fa-2x');
                );
                if ( obj.criticality === 'info')
                {
                    feed_li.append(
                        a_element = $('<a>').attr('href','http://' + document.domain + '/detail/'+obj.hash);
                    );
                }
                else
                {
                    feed_li.append(
                        a_element = $('<a>').attr('href','http://' + document.domain + '/detail/'+obj.hash, 'class', 'strain');
                    );
                }
                a_element = $('<p>').attr('class', 'feed_title').append(obj.title);
                a_element = $('<p>').append(obj.msg);
                )));
            });
        };
</script>
