const bcrypt = require('bcrypt');
const { conference: ConferenceModels } = require('../models');
require('dotenv').config();

// const status = ["reject", "pendding", "approve"]

const View = async (req, res) => {
    try {
        const request = req.body
        const user_id = req.user.user_id
        const find = { } 
        if (req.user.user_role == "user") {
            if (user_id) {
                find.$or = [
                  {
                    conference_status: 2,
                    conference_created_by: { $ne: user_id },
                  },
                  {
                    conference_created_by: user_id,
                  },
                ];
            } else {
               find.conference_status = 2
               conference_deleted_at: null
            }
        }
        if(request.conference_id) find._id = request.conference_id
        if(request.conferece_kapasitas) find.conferece_kapasitas = request.conferece_kapasitas
        if(request.conferece_status) find.conferece_status = request.conferece_status
        const conference = await ConferenceModels.find(find).sort({ conference_nama: 1 }).select("conference_nama conference_created_by conference_kapasitas conference_status conference_waktu")
        return res.json({
            error: false,
            code: 200,
            message: 'Data Berhasil Ditemukan Berhasil',
            description: 'Berhasil menemukan data!',
            data: conference,
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
        if (!request.conference_nama || !request.conference_start || !request.conference_end || !request.conference_kapasitas) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses!"
            });
        } 
        const conferenceFind = await ConferenceModels.findOne({
            conference_created_by: req.user.user_id, // Dibuat oleh user yang sama
            $or: [
                {
                    conference_nama: request.conference_nama, // Nama tidak boleh sama
                },
                {
                    $or: [
                        // Memeriksa apakah waktu mulai dan akhir tumpang tindih
                        {
                            "conference_waktu.start": { $lte: new Date(request.conference_end) },
                            "conference_waktu.end": { $gte: new Date(request.conference_start) },
                        },
                        // Memeriksa apakah konferensi baru berada dalam rentang waktu konferensi lama
                        {
                            "conference_waktu.start": { $gte: new Date(request.conference_start) },
                            "conference_waktu.end": { $lte: new Date(request.conference_end) },
                        },
                        // Memeriksa apakah waktu mulai dan akhir baru tumpang tindih
                        {
                            "conference_waktu.start": { $gte: new Date(request.conference_start) },
                            "conference_waktu.end": { $gte: new Date(request.conference_end) },
                        },
                    ],
                },
                {
                    // Jika nama berbeda, hanya periksa apakah waktu bertabrakan
                    conference_nama: { $ne: request.conference_nama },
                    $or: [
                        {
                            "conference_waktu.start": { $lte: new Date(request.conference_end) },
                            "conference_waktu.end": { $gte: new Date(request.conference_start) },
                        },
                    ],
                },
            ],
        });

        if (conferenceFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: "Bad Request",
                description: "Nama sudah digunakan atau periode waktu tumpang tindih!",
            });
        }

        if (conferenceFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: 'Nama atau periode atau akses tidak sesuai!',
            });
        }

        const conference = new ConferenceModels({
            conference_nama: request.conference_nama, 
            conference_deskripsi: request.conference_deskripsi, 
            conference_kapasitas: request.conference_kapasitas, 
            conference_waktu: {
                start: new Date(request.conference_start),
                end: new Date(request.conference_end),
            }, 
            conference_created_by: req.user.user_id, 
            conference_created_at: new Date()
        })
        await conference.save()
        return res.json({
            error: false,
            code: 200,
            message: 'Success',
            description: 'Berhasil menyimpan data!',
            data: conference
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

const Update = async (req, res) => {
    try {
        const {conference_status, ...request} = req.body;
        let condition = {
            _id: { $ne: request.conference_id }, // Memastikan ID berbeda
            conference_created_by: req.user.user_id, // Dibuat oleh user yang sama
            conference_nama: request.conference_nama,
        };
        if (request.conference_start && request.conference_end) {
            condition = {
                _id: { $ne: request.conference_id }, // Memastikan ID berbeda
                conference_created_by: req.user.user_id, // Dibuat oleh user yang sama
                $or: [
                    {
                        conference_nama: request.conference_nama,
                    },
                    // Memeriksa apakah waktu mulai dan akhir tumpang tindih
                    {
                        "conference_waktu.start": { $lte: new Date(request.conference_end) },
                        "conference_waktu.end": { $gte: new Date(request.conference_start) },
                    },
                    // Memeriksa apakah konferensi baru berada dalam rentang waktu konferensi lama
                    {
                        "conference_waktu.start": { $gte: new Date(request.conference_start) },
                        "conference_waktu.end": { $lte: new Date(request.conference_end) },
                    },
                    // Memeriksa apakah waktu mulai dan akhir baru tumpang tindih
                    {
                        "conference_waktu.start": { $gte: new Date(request.conference_start) },
                        "conference_waktu.end": { $gte: new Date(request.conference_end) },
                    },
                    {
                        // Jika nama berbeda, hanya periksa apakah waktu bertabrakan
                        conference_nama: { $ne: request.conference_nama },
                        $or: [
                            {
                                "conference_waktu.start": { $lte: new Date(request.conference_end) },
                                "conference_waktu.end": { $gte: new Date(request.conference_start) },
                            },
                        ],
                    },
                ]
            };
        }

        const roleCondition = req.user.user_role === "user" ? { conference_created_by: req.user.user_id } : {};
        const conflictFind = await ConferenceModels.find({
            ...condition,
            ...roleCondition
        });

        if (conflictFind.length > 0) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: "Bad Request",
                description: "Nama sudah digunakan atau periode waktu tumpang tindih!",
            });
        }

        const aksesRequest = req.user.user_role == "user" ? { 
           ...request,
           conference_waktu: {
                start: request.conference_start ? new Date(request.conference_start) : undefined, // Gunakan tanggal lama jika tidak ada inputan
                end: request.conference_end ? new Date(request.conference_end) : undefined,       // Gunakan tanggal lama jika tidak ada inputan
            },
        } : { conference_status: conference_status }
        const conference = await ConferenceModels.findByIdAndUpdate(
            request.conference_id,
            {
                $set: {
                    ...aksesRequest,
                    conference_updated_at: new Date(),
                },
            },
            { new: true, runValidators: true }
        );
        if (!conference) {
            return res.status(404).json({
                error: true,
                code: 404,
                message: 'Not Found',
                description: 'Data tidak ditemukan.',
            });
        }
        return res.json({
            error: false,
            code: 200,
            message: 'Update Berhasil',
            description: 'Berhasil melakukan update!',
            data: conference,
        });
    } catch (error) {
        return res.status(500).json({
            error: true,
            code: 500,
            message: 'Terjadi Kesalahan Server',
            description: 'Permintaan tidak dapat diproses karena terjadi kesalahan pada server.',
            errorMessage: error.message,
        });
    }
}

const Delete = async (req, res) => {
    try {
        const { conference_id } = req.body
        let permanent = ""
        if (!conference_id) {
            return res.status(400).json({
                code: 400,
                message: 'Bad Request',
                description: "Pilih data yang ingin dihapus!"
            })
        }
        let params = {}
        if(req.user.user_role == "user") {
            params = {
                $and: [
                    { _id: conference_id },
                    { conference_created_by: req.user.user_id },
                ],
            }
        } else {
            params = { _id: conference_id }
        }
        const conferenceFind = await ConferenceModels.findOne(params)
        if(!conferenceFind) {
            return res.status(403).json({
                error: true,
                code: 403,
                message: 'Forbidden',
                description: 'Anda tidak memiliki akses ke resource ini.',
            })
        }
        if (conferenceFind.conference_deleted_at) {
            await ConferenceModels.deleteOne(
                {_id: conference_id}
            )
            permanent = " Permanen"
        } else {
            const conferenceDelete = await ConferenceModels.findOneAndUpdate(
                { _id: conference_id }, 
                {
                    conference_updated_at: new Date(),
                    conference_deleted_at: new Date()
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
                conference_id: conference_id
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
    View, Store, Update, Delete
}