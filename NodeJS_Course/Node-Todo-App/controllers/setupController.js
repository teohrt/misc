var Todos = require('../models/todoModels');

module.exports = function(app) {

    app.get('/api/setupTodos', function(req, res) {

        // Seed data
        var starterTodos = [
            {
                username: 'tyypo',
                todo: 'correct spelling',
                isDone: false,
                hasAttachment: false
            },
            {
                username: 'Dorothy Joke',
                todo: 'We\'re not in Kansas anymore...todo',
                isDone: true,
                hasAttachment: false
            },
            {
                username: 'Trace',
                todo: 'learn node',
                isDone: false,
                hasAttachment: false
            }
        ];
        // Add data into database
        Todos.create(starterTodos, function(err, results) {
            res.send(results);
            if (err) {
                console.log(err);
                return ;
            }
        });
    });

}