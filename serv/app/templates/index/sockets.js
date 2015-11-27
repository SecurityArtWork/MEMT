<script type="text/javascript" charset="utf-8">
        $(document).ready(function(){
            RTMapNamespace = '/rtmap'; // change to an empty string to use the global namespace

            // the socket.io documentation recommends sending an explicit package upon connection
            // this is specially important when using the global namespace
            var rtsocket = io.connect('http://' + document.domain + ':' + location.port + RTMapNamespace);

            // New attack.
            socket.on('update', function(data) {

            });

            // event handler for new connections
            socket.on('connect', function(data) {

            });
        });
</script>
