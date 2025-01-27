const mongoose = require('mongoose');

const schema = new mongoose.Schema(
    {
        conference_nama: { type: String, default: null, required: true }, 
        conference_deskripsi: { type: String, default: null }, 
        conference_waktu: { 
            type: Object,
            default: {
                start: { type: Date, default: null }, 
                end: { type: Date, default: null }, 
            }
        }, 
        conference_kapasitas: { type: Number, default: 0 },
        conference_status: { type: Number, default: 1 },
        conference_created_by: { 
            type: Object, 
            default: {
                id: { type: String, default: null },        
                nama: { type: String, default: null },      
                username: { type: String, default: null },  
                email: { type: String, default: null }   
            } 
        },
        conference_created_at: { type: Date, default: Date.now },
        conference_updated_at: { type: Date, default: null },
        conference_deleted_at: { type: Date, default: null },
    }
);

module.exports = mongoose.model("conference", schema);
