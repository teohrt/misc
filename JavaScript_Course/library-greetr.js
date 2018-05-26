(function(global, $) {

    // 'new' object
    var Greetr = function(firstname, lastname, language) {
        return new Greetr.init(firstname, lastname, language);
    }

    // hidden inside the immediately invoked function expression.
    // not directly accessible
    var supportedLangs = ['en', 'es'];

    // informa greetings
    var greetings = {
        en: 'Hello',
        es: 'Hola'
    };

    // formal greetings
    var formalGreetings = {
        en: 'Greetings',
        es: 'Saludos'
    };

    // logger messages
    var logMessages = {
        en: 'Logged in',
        es: 'Inicio sesion'
    };

    // prototype holds methods to save memory
    Greetr.prototype = {

        // this refers to the calling object at excecution
        fullname: function() {
            return this.firstname + ' ' + this.lastname;
        },

        // checks that the language is valid
        // uses the externally inaccessible 'supportedLangs'
        validate: function() {
            
             if ( supportedLangs.indexOf(this.language) === -1 ) {
                throw "Invalid Language";
             }
        },

        // retrieve message by using [] syntax
        greeting: function() {
            return greetings[this.language] + ' ' +this.firstname + '!';
        },

        formalGreeting: function() {
            return formalGreetings[this.language] + ', ' + this.fullname() + '.'; 
        },

        // the chainable methods return their own containing object
        greet: function(formal) {
            var msg;

            //if undefined or null it will be coerced to 'false'
            if (formal) {
                msg = this.formalGreeting();
            }            
            else {
                msg = this.greeting();
            }

            if (console) {
                console.log(msg);
            }

            //'this' refer to the calling object at execution time
            //makes the method chainable
            return this;
        },

        log: function() {
            if (console) {
                console.log(logMessages[this.language] + ': ' + this.fullname());
            }
            return this;
        },

        setLang: function(lang) {

            // set the language
            this.language = lang;

            // then validate
            this.validate();

            //make chainable
            return this;
        },

        HTMLGreeting: function(selector, formal) {

            // checks jQuery
            if (!$) {
                throw 'jQuery not loaded';
            }
            else {
                msg = this.greeting();
            }

            $(selector).html(msg);

            return this;
            //making the code chainable!
        }

    };

    // the object is created here, so we can 'new' and object without using the keyword 'new'
    Greetr.init = function(firstname, lastname, language) {

        var self = this;
        self.firstname = firstname || '';
        self.lastname = lastname || '';
        self.language = language || 'en';

        self.validate();
    }

    // this is the trick borrowed from jQuery so we don't have to use the 'new' keyword
    Greetr.init.prototype = Greetr.prototype;

    // attach out Greetr to the global object, and orovide a shorthand 'G$' for ease of use
    global.Greetr = global.G$ = Greetr;

}(window, jQuery));