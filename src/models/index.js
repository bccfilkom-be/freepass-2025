const mongoose = require('mongoose');

module.exports = {
    user: require('./usermodel.js'),
    conference: require('./conferenceModel.js'),
    feedback: require('./feedbackModel.js'),
    conferenceuserjoin: require('./conferenceuserjoinModel.js'),
}