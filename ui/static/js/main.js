$(function () {
    var socketHost = '/';

    $('#add-user').submit(
        function (e) {
            e.preventDefault();

            var attributes = JSON.parse($('#attributes').val()),
                callbackUrl = $('#callbackUrl').val(),
                id = Math.random().toString(36).substring(7),
                user = {id, attributes, callbackUrl}

            $.ajax(
                {
                    url: '/user/encrypt',
                    dataType: 'json',
                    contentType: 'application/json',
                    async: false,
                    type: 'POST',
                    data: JSON.stringify(user)
                }
            ).then(data => {
                var options = data.options,
                    socket = io(socketHost, options);

                var emit = function (event, message) {

                    $('#ul-' + id).append(
                        '<li class="list-group-item list-group-item-success" > ' + event + ' : ' + JSON.stringify(message) + '</li>'
                    )

                    socket.emit(event, message, response => {
                        $('#ul-' + id).append(
                            '<li class="list-group-item" > ' + event + ' : ' + JSON.stringify(response) + '</li>'
                        )
                    })
                }

                var onevent = socket.onevent;
                socket.onevent = function (packet) {
                    var args = packet.data || [];
                    onevent.call (this, packet);    // original call
                    packet.data = ["*"].concat(args);
                    onevent.call(this, packet);      // additional call to catch-all
                };

                options.forceRecreate = true;

                socket.on('connect', function () {
                    setTimeout(function () {
                        emit('auth', data);
                    }, 10);
                });

                socket.on('*', function (event, message, ack) {
                    $('#ul-' + id).append(
                        '<li class="list-group-item" > ' + event + ' : ' + JSON.stringify(message) + '</li>'
                    )

                    typeof ack === 'function' ? ack() : null
                });

                $('#join').on('click', function () {
                        var realm = $('#realm').val(),
                            room = $('#room').val()

                        emit('room/join', {realm, room})
                    }
                )

                $('#leave').on('click', function () {
                        var realm = $('#realm').val(),
                            room = $('#room').val(),
                            userId = $('#userId').val()

                        emit('room/leave', {realm, room})
                    }
                )

                $('#info').on('click', function () {
                        var realm = $('#realm').val()

                        emit('room/info', {realm})
                    }
                )

                $('#' + id + '-disconnect').on('click', function (e) {
                    e.preventDefault();
                    socket.disconnect();
                    $('#' + id).remove();
                });
            });

            var block = '<div class="card card-body" id="' + id + '"><h3>' + id + '<button id="' + id + '-disconnect" class="btn btn-danger">X</button></h3>' + JSON.stringify(user) + '<ul class="list-group user-log" id="ul-' + id + '"></ul></div>';
            $('#chat').append(block);
            $('#username').val('');

        }
    );

    $('#send-message').on('submit', function (e) {
        e.preventDefault();

        let data = {payload:JSON.parse( $('#payload').val()), query: JSON.parse($('#query').val())};
        $.ajax(
            {
                url: '/notification/send',
                dataType: 'json',
                contentType: 'application/json',
                async: false,
                type: 'POST',
                data: JSON.stringify(data)
            }
        );
    });

    $('#clear-log').on('click', function (e) {
        e.preventDefault();
        $('.user-log').empty();
    });
});