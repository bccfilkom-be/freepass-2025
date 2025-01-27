const mongoose = require('mongoose');

const schema = new mongoose.Schema(
    {
        feedback_conference: { 
            type: Object, 
            default: {
                id: { type: String, default: null },
                nama: { type: String, default: null },
            } 
        }, 
        feedback_user: { 
            id: { type: String, default: null },        
            email: { type: String, default: null }    
        },
        feedback_rating: { type: Number, default: 0 },
        feedback_message: { type: String, default: null },
        feedback_created_at: { type: Date, default: Date.now },
        feedback_updated_at: { type: Date, default: null },
        feedback_deleted_at: { type: Date, default: null },
    }
);

module.exports = mongoose.model("feedback", schema);
