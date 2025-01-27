const bcrypt = require('bcrypt');
const mongoose = require("mongoose");
const { user: UserModels, conference: ConferenceModels,  conferenceuserjoin: ConferenceuserjoinModels} = require('../models');
require('dotenv').config();

const View = async (req, res) => {
    try {
        const request = req.body
        const user_id = req.user.user_id
        const find = { 
            conferenceuserjoin_deleted_at: null,
            "conferenceuserjoin_user.email": req.user.user_email
        } 
        console.log(find)
        const conferencejoin = await ConferenceuserjoinModels.find(find).sort({ conferenceuserjoin_created_at: 1 })
        return res.json({
            error: false,
            code: 200,
            message: 'Data Berhasil Ditemukan Berhasil',
            description: 'Berhasil menemukan data!',
            data: conferencejoin,
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

        const userFind = await UserModels.findOne({_id: req.user.user_id});
        const conferenceFind = await ConferenceModels.findOne({
            _id: request.conference_id, 
            conference_created_by:  { $ne: req.user.user_id }, 
            conference_status: 2,
        });
        if (!request.conference_id || !conferenceFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses!"
            });
        } 
        console.log(conferenceFind)
        let maxConference = conferenceFind.conference_kapasitas
        const conferencejoinFind = await ConferenceuserjoinModels.findOne({
            "conferenceuserjoin_user.id": userFind._id,
            $or: [
                {
                    $or: [
                        // Memeriksa apakah waktu mulai dan akhir tumpang tindih
                        {
                            "conferenceuserjoin_conference.time.start": { $lte: new Date(conferenceFind.conference_waktu.end) },
                            "conferenceuserjoin_conference.time.end": { $gte: new Date(conferenceFind.conference_waktu.start) },
                        },
                        // Memeriksa apakah konferensi baru berada dalam rentang waktu konferensi lama
                        {
                            "conferenceuserjoin_conference.time.start": { $gte: new Date(conferenceFind.conference_waktu.start) },
                            "conferenceuserjoin_conference.time.end": { $lte: new Date(conferenceFind.conference_waktu.end) },
                        },
                        // Memeriksa apakah waktu mulai dan akhir baru tumpang tindih
                        {
                            "conferenceuserjoin_conference.time.start": { $gte: new Date(conferenceFind.conference_waktu.start) },
                            "conferenceuserjoin_conference.time.end": { $gte: new Date(conferenceFind.conference_waktu.end) },
                        },
                    ],
                },
                {
                    // periksa apakah waktu bertabrakan
                    $or: [
                        {
                            "conferenceuserjoin_conference.time.start": { $lte: new Date(conferenceFind.conference_waktu.end) },
                            "conferenceuserjoin_conference.time.end": { $gte: new Date(conferenceFind.conference_waktu.start) },
                        },
                    ],
                },
            ],
        })
        console.log(conferencejoinFind)
        if (conferencejoinFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Hanya bisa memilih 1 kali dalam periode yang sama!"
            });
        } 
        const limitUser = await ConferenceuserjoinModels.find({"conferenceuserjoin_conference.id": request.conference_id})
        if (limitUser.length >= maxConference) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Kuota sudah terpenuhi!"
            });
        } 
        const conferenceuserjoin = new ConferenceuserjoinModels({
            conferenceuserjoin_conference: {
                id: request.conference_id, 
                nama: conferenceFind.conference_nama, 
                time: {
                    start: conferenceFind.conference_waktu.start,
                    end: conferenceFind.conference_waktu.end,
                }, 
            },
            conferenceuserjoin_user: {
                id: userFind._id, 
                email: userFind.user_email
            }, 
            conferenceuserjoin_created_at: new Date()
        })
        await conferenceuserjoin.save()
        return res.json({
            error: false,
            code: 200,
            message: 'Success',
            description: 'Berhasil menyimpan data!',
            data: conferenceuserjoin
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
        const { conferencejoinuser_id } = req.body
        let permanent = ""
        if (!conferencejoinuser_id) {
            return res.status(400).json({
                code: 400,
                message: 'Bad Request',
                description: "Pilih data yang ingin dihapus!"
            })
        }
        const conferenceuserFind = await ConferenceuserjoinModels.findOne({
            _id: conferencejoinuser_id
        })
        if (conferenceuserFind.conferenceuserjoin_deleted_at) {
            await ConferenceModels.deleteOne(
                {_id: conferencejoinuser_id}
            )
            permanent = " Permanen"
        } else {
            const conferenceDelete = await ConferenceModels.findOneAndUpdate(
                { _id: conferencejoinuser_id }, 
                {
                    conferenceuserjoin_updated_at: new Date(),
                    conferenceuserjoin_deleted_at: new Date()
                },
                { new: true },
            )
        }
        return  res.status(200).json({
            error: false,
            code: 200,
            message: `Data Berhasil Dihapus${permanent}`,
            description: 'Berhasil menghapus data',
            data: {
                conferencejoinuser_id: conferencejoinuser_id
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