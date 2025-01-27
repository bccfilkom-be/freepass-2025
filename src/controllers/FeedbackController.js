const bcrypt = require('bcrypt');
const { user: UserModels, conference: ConferenceModels, feedback: FeedbackModels, conferenceuserjoin: ConferenceuserjoinModels } = require('../models');
require('dotenv').config();

const View = async (req, res) => {
    try {
        const request = req.body
        const user = req.user
        const find = {} 
        if(user.user_role == "user") {
            find["feedback_user.id"] = user.user_id
        }
        if(request.conference_id) find["feedback_conference.id"] = request.conference_id
        if(request.feedback_rating) find["feedback_rating"] = request.feedback_rating
        
        const feedback = await FeedbackModels.find(find).sort({ feedback_created_at: 1 })
        return res.json({
            error: false,
            code: 200,
            message: 'Data Berhasil Ditemukan Berhasil',
            description: 'Berhasil menemukan data!',
            data: feedback,
        })
    } catch (error) {
        return res.status(500).json({
            error: true,
            code: 500,
            message: 'Terjadi Kesalahan Server',
            description: 'Permintaan tidak dapat diproses karena terjadi kesalahan pada server.',
            errorMessage: error.message
        })
    }
}

const Store = async (req, res) => {
    try {
        const request = req.body
        const userFind = await UserModels.findOne({_id: req.user.user_id})
        const feedbackFind = await FeedbackModels.findOne({
            "feedback_user.id": req.user.user_id,
            "feedback_conference.id": request.conference_id
        })
        const conferenceFind = await ConferenceModels.findOne({
            _id: request.conference_id, 
            conference_created_by:  { $ne: req.user.user_id }, 
            conference_status: 2,
        });
        if (!request.conference_id || !conferenceFind || !request.feedback_rating || request.feedback_rating >= 5 || !request.feedback_message || feedbackFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses!"
            });
        } 
        const conferencejoinFind = await ConferenceuserjoinModels.findOne({
            "conferenceuserjoin_conference.id": request.conference_id,
            "conferenceuserjoin_user.id": userFind._id
        })
        if (!conferencejoinFind) {
            return res.status(400).json({
                error: true,
                code: 404,
                message: 'Not Found',
                description: "Data pertemuan anda tidak tersedia!"
            });
        } 
        const feedback = new FeedbackModels({
            feedback_conference: {
                id: request.conference_id, 
                nama: conferenceFind.conference_nama, 
            },
            feedback_user: {
                id: userFind._id, 
                nama: userFind.user_nama,
                username: userFind.user_username,
                email: userFind.user_email
            }, 
            feedback_rating: request.feedback_rating,
            feedback_message: request.feedback_message,
            feedback_created_at: new Date()
        })
        await feedback.save()
        return res.json({
            error: false,
            code: 200,
            message: 'Success',
            description: 'Berhasil menyimpan data!',
            data: feedback
        })
    } catch (error) {
        return res.status(500).json({
            error: true,
            code: 500,
            message: 'Terjadi Kesalahan Server',
            description: 'Permintaan tidak dapat diproses karena terjadi kesalahan pada server.',
            errorMessage: error.message
        })
    }
}

const Delete = async (req, res) => {
    try {
        const { feedback_id } = req.body
        if (!feedback_id) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Pilih data yang akan dihapus!"
            });
        } 
        let permanent = ""
        const feedbackFind = await FeedbackModels.findOne({
            _id: feedback_id
        })
        if (feedbackFind.feedback_deleted_at) {
            await FeedbackModels.deleteOne(
                {_id: feedback_id}
            )
            permanent = " Permanen"
        } else {
            const feedbackDelete = await FeedbackModels.findOneAndUpdate(
                { _id: feedback_id }, 
                {
                    feedback_updated_at: new Date(),
                    feedback_deleted_at: new Date()
                },
                { new: true },
            )
        }
        return res.status(200).json({
            error: false,
            code: 200,
            message: `Data Berhasil Dihapus${permanent}`,
            description: 'Berhasil menghapus data',
            data: {
                feedback_id: feedback_id
            }
        })
    } catch (error) {
        return res.status(500).json({
            error: true,
            code: 500,
            message: 'Terjadi Kesalahan Server',
            description: 'Permintaan tidak dapat diproses karena terjadi kesalahan pada server.',
            errorMessage: error.message
        })
    }
}

module.exports = {
    View, Store, Delete
}