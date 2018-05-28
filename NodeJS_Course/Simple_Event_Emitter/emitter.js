function Emitter() {
    this.events = {};
}

Emitter.prototype.on = function(type, listener) {
    // if the property already exists, alright. otherwise make a new array
    this.events[type] = this.events[type] || [];
    this.events[type].push(listener);
}

Emitter.prototype.emit = function(type) {
    // if the type exists then loop thu the type's array
    // and invoke all its listener functions
    if (this.events[type]) {
        this.events[type].forEach(function(listener) {
            listener();
        });
    }
}

// make the constructor availalble
// this way you can create many Emitters at once
module.exports = Emitter;