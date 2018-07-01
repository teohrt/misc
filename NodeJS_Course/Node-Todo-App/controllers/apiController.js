var Todos = require('../models/todoModels');
var bodyParser = require('body-parser');
var ObjectID = require('mongodb').ObjectID;

module.exports = function(app) {

    app.use(bodyParser.json());
    app.use(bodyParser.urlencoded({ extended: true}));

    // Get all Todos for a person
    app.get('/api/todos/:uname', function(req, res) {
        Todos.find({ username: req.params.uname }, function(err, todos) {
            if (err) throw err;

            res.send(todos);
        })
    })

    // Get a particular Todo by its ID
    app.get('/api/todo/:id', function(req, res) {
        Todos.findById(new ObjectID(req.params.id), function(err, todo) {
            if (err) throw err;

            res.send(todo);
        });
    });

    // Posting a Todo. Just JSON within an http request
    app.post('/api/todo', function(req, res) {

        // Updates if already present
        if(req.body._id) {
            Todos.findByIdAndUpdate(new ObjectID(req.body._id), { todo: req.body.todo, isDone: req.body.isDone, hasAttachment: req.body.hasAttachment }, function(err, todo) {
                if (err) throw err;

                res.send('Updated record');
            });
        }

        // Create new Todo and put on db
        else {
            var newTodo = Todos({ 
                username: 'test',
                todo: req.body.todo,
                isDone: req.body.isDone,
                hasAttachment: req.body.hasAttachment
            });

            newTodo.save(function(err) {
                if (err) throw err;
                res.send('No match. Added new record');
            });
        }

    });

    // Remove a todo
    app.delete('/api/todo', function(req, res) {

        Todos.findByIdAndRemove(new ObjectID(req.body._id), function(err) {
            if (err) throw err;
            res.send('Success');
        })

    });

}