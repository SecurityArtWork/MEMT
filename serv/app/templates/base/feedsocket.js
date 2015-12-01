<script type="text/javascript" charset="utf-8">
        $(document).ready(function(){
            feedNamespace = '/feed'; // change to an empty string to use the global namespace

            // the socket.io documentation recommends sending an explicit package upon connection
            // this is specially important when using the global namespace
            var feedsocket = io.connect('http://' + document.domain + ':' + location.port + feedNamespace);

            // event handler for server sent data
            // the data is displayed in the "Received" section of the page
            feedsocket.on('update', function(feeds) {
                parseFeedData(feeds);
            });

            // event handler for new connections
            feedsocket.on('connect', function(feeds) {
                parseFeedData(feeds);
            });
        });

        function parseFeedData(feeds){
            if(feeds != undefined){
                feeds = JSON.parse(feeds)
                console.dir(feeds);
                feed_ul = $('#feed');
                $.each(feeds, function(idx, obj) {
                    console.log(obj);
                    feed_li = feed_ul.append($('<li>'));
                    console.log(feed_li);
                    feed_li.append($('<i>').attr('class', 'fa fa-fw fa-bullhorn fa-2x'));
                    a_element = $('<a>').attr('href','http://' + document.domain + '/detail/'+obj.sha256)
                    a_element = $('<p>').attr('class', 'feed_title').append(obj.title[0]);
                    a_element = $('<p>').append(obj.msg);
                    if ( obj.criticaly[0] == 'danger') {
                        feed_li.append(a_element.attr('class', 'strain'));
                    }
                    feed_li.append(a_element);
                });
            }
        };
</script>
