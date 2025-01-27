const express = require("express");
const { ConferenJoinController, FeedbackController } = require('./../controllers');

const User = express.Router();

User.get('/', (req, res) => {
    return res.json({
        message: 'Welcome, User!',
        user: req.user,
    });
});

User.get('/join', ConferenJoinController.View);
User.post('/join', ConferenJoinController.Store);
User.delete('/join', ConferenJoinController.Delete);
User.get('/feedback', FeedbackController.View);
User.post('/feedback', FeedbackController.Store);

module.exports = User