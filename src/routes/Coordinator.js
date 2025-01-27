const express = require("express");
const { ConferenceController, FeedbackController } = require('../controllers');

const Coordinator = express.Router();

Coordinator.get('/', (req, res) => {
    return res.json({
        message: 'Welcome, Coordinator!',
        user: req.user,
    });
});
Coordinator.get('/feedback', FeedbackController.View)
Coordinator.delete('/feedback', FeedbackController.Delete)
module.exports = Coordinator