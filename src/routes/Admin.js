const express = require("express");
const { ProfileController } = require('../controllers');

const Admin = express.Router();

Admin.get('/', (req, res) => {
    return res.json({
        message: 'Welcome, Admin!',
        user: req.user,
    });
});
Admin.post("/user", ProfileController.Store)
Admin.patch("/user", ProfileController.UpdateDataUser)
Admin.delete("/user", ProfileController.Delete)

module.exports = Admin