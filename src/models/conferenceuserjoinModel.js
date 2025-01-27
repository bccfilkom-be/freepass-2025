const mongoose = require('mongoose');

const schema = new mongoose.Schema(
    {
        // conferenceuserjoin_conference_id: { type: String, default: null }, 
        conferenceuserjoin_conference: { 
            type: Object, 
            default: {
                id: { type: String, default: null },
                nama: { type: String, default: null },
                time: {
                    start: { type: Date, default: null },
                    end: { type: Date, default: null },
                }
            } 
        }, 
        conferenceuserjoin_user: {
            type: Object,
            default: {
                id: { type: String, default: null },        
                nama: { type: String, default: null },      
                username: { type: String, default: null },  
                email: { type: String, default: null }      
            },
        },
        conferenceuserjoin_created_at: { type: Date, default: Date.now },
        conferenceuserjoin_updated_at: { type: Date, default: null },
        conferenceuserjoin_deleted_at: { type: Date, default: null },
    }
);

module.exports = mongoose.model("conferenceuserjoin", schema);
