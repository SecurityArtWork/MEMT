<script type="text/javascript" charset="utf-8">
        $(document).ready(function(){

            celery = '/celery';

            // the socket.io documentation recommends sending an explicit package upon connection
            // this is specially important when using the global namespace
            var celerysocket = io.connect('http://' + document.domain + ':' + location.port + celery);

            // New sample received from
            rtsocket.on('update', function(data) {
                if data['status'] == 200{
                    redirect(data['hash']);
                }
            });

        });
        //Let's parse the JSON data for each item
        function redirect(hash) {
            window.location.href = 'http://' + document.domain + '/detail/'+ hash;
        };
</script>
