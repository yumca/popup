var wsSocket = {
    conn: null,
    connect: function (url) {
        var _this = this
        this.conn = new WebSocket(url)
        this.conn.onopen = function (event) {
            // this.data = '连接成功'
            $('#contents').append("<div>" + url + "连接成功</div>");
        };
        this.conn.onmessage = function (event) {
            data = JSON.parse(event.data)
            console.log('onmessage', event.data)
            if (_this.callbacks[data.signal] != undefined) {
                _this.callbacks[data.signal](data,_this)
            }
            // $('#contents').append("<div>" + event.data + "</div>");
        };
        this.conn.onclose = function (event) {
            // this.data = '连接已关闭'
            $('#contents').append("<div>" + url + "连接已关闭</div>");
        };
        return this
    },
    data: {
        user:{
            online:false,
            id:0,
            uname:'',
            fd:0,
            loginkey:''
        },
        game:{

        }
    },
    callbacks: {},
    sendone: function (tfd,data,callback) {
        sig = {
            signal: 'sendone',
            fd: user.wsfd,
            tfd: tfd,
            loginkey:user.loginkey,
            data: {
                message: data
            }
        }
        this.conn.send(JSON.stringify(sig));
        this.callbacks['sendone'] = callback
    },
    sendall: function (data) {
        sig = {
            signal: 'sendall',
            fd: user.wsfd,
            loginkey:user.loginkey,
            data: {
                message: data
            }
        }
        this.conn.send(JSON.stringify(sig));
        this.callbacks['sendall'] = callback
    },
    bindfdwithloginkey: function (user, callback) {
        sig = {
            signal: 'bindfdwithloginkey',
            fd: user.wsfd,
            loginkey:user.loginkey,
            data: {}
        }
        this.conn.send(JSON.stringify(sig))
        this.callbacks['bindfdwithloginkey'] = callback
    }
}
