const mongoose = require('mongoose');

const schema = new mongoose.Schema(
    {
        user_nama: { type: String, default: null }, 
        user_email: { type: String, default: null },
        user_username: { type: String, default: null },
        user_role: { type: String, default: "user" },
        user_password: { type: String, default: null },
        user_token: { type: String, default: null },
        user_created_at: { type: Date, default: Date.now },
        user_updated_at: { type: Date, default: null },
        user_deleted_at: { type: Date, default: null },
    }
);

module.exports = mongoose.model("users", schema);
