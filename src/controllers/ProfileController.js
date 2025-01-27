const bcrypt = require('bcrypt');
const validator = require('validator');
const { user: UserModels, } = require('../models');
require('dotenv').config();

const allowedRoles = ['admin', 'user', 'coordinator'];

const Search = async (req, res) => {
    try {
        const request = req.body
        const user_role = req.user.user_role
        const find = { 
            user_deleted_at: null,
        } 
        if(request.user_id) find._id = request.user_id
        if(request.user_username) find.user_username = request.user_username
        if(user_role == "user") find.user_role= user_role
        const user = await UserModels.find(find).sort({ user_email: -1 }).select("user_username user_email user_nama user_role")
        return res.json({
            error: false,
            code: 200,
            message: 'Data Berhasil Ditemukan Berhasil',
            description: 'Berhasil menemukan data!',
            data: user,
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
        const userFind = await UserModels.findOne({
            user_email: request.user_email,
            $or: [
                { user_username: request.user_username },
            ]
        })
       
        if (userFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Data sudah digunakan!"
            });
        }
        if (!allowedRoles.includes(request.user_role)) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Data role tidak tersedia!"
            });
        }
        if (!request.user_email || !validator.isEmail(request.user_email)) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses!"
            });
        } 

        const hashedPassword = await bcrypt.hash("password123", 10)
        const payload = { user_role: request.user_role, user_email: request.user_email, user_password: hashedPassword, user_created_at: new Date() }
        
        const user = new UserModels(payload)
        await user.save()
        return res.json({
            error: false,
            code: 200,
            message: 'Success',
            description: 'Berhasil menyimpan data!',
            data: userFind
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
        const request = req.body
        const userId = req.user.user_id
        if (request.user_email) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses!"
            });
        } 
        
        const user = await UserModels.findByIdAndUpdate(
            userId,
            {
                $set: { 
                    ...request, 
                    user_role: req.user.user_role,
                    user_updated_at: new Date() 
                },
            },
            { new: true, runValidators: true }
        );
        
        return res.json({
            error: false,
            code: 200,
            message: 'Update Berhasil',
            description: 'Berhasil melakukan update!',
            data: user
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

const UpdateDataUser = async (req, res) => {
    try {
        const request = req.body
        if (request.user_email) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses!"
            });
        } 
        
        const user = await UserModels.findByIdAndUpdate(
            user_id,
            {
                $set: { 
                    ...request,
                    user_role: request.user_role,
                    user_updated_at: new Date() 
                },
            },
            { new: true, runValidators: true }
        );
        
        return res.json({
            error: false,
            code: 200,
            message: 'Update Berhasil',
            description: 'Berhasil melakukan update!',
            data: user
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
        const { user_id } = req.body
        let permanent = ""
        if (!user_id) {
            return res.status(400).json({
                code: 400,
                message: 'Bad Request',
                description: "Pilih data yang ingin dihapus!"
            })
        }
        const userFind = await UserModels.findOne({_id: user_id})
        if (userFind.user_deleted_at) {
            await UserModels.deleteOne(
                {_id: user_id}
            )
            permanent = "Permanen"
        } else {
            const userDelete = await UserModels.findOneAndUpdate(
                { _id: user_id }, 
                {
                    user_updated_at: new Date(),
                    user_deleted_at: new Date()
                },
                { new: true },
            )
        }
        return  res.status(200).json({
            error: false,
            code: 200,
            message: `Data Berhasil Dihapus ${permanent}`,
            description: 'Berhasil menghapus data',
            data: {
                user_id: user_id
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
    Update, Search, Store, UpdateDataUser, Delete
}