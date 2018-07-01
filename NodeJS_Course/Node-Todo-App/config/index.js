// We have the username and password as an object available to us
// No encryption because laziness
var configValues = require('./config');

// Returns database URL for connection
module.exports = {

    getDbConnectionString: function() {
        //return 'mongodb://' + configValues.uname + ':' + configValues.pwd + '@ds217671.mlab.com:17671//hawtstorage'
        // this kind of auth doesn't work for that service anymore
        // going local
        return 'mongodb://localhost/todo-app'
    }

}