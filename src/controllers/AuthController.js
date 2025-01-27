const bcrypt = require('bcrypt');
const { user: UserModels } = require('./../models');
const validator = require('validator');
const jwt = require('jsonwebtoken');
require('dotenv').config();

const Register = async (req, res) => {
    try {
        const request = req.body

        const userFind = await UserModels.findOne({
            user_email: request.user_email,
            $or: [
                { user_username: request.user_username },
            ]
        })
        if (!request.user_email || !validator.isEmail(request.user_email) || !request.user_password || !request.user_confirmPassword) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses karena ada data yang hilang atau tidak valid!"
            });
        } 
        if (userFind) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Email sudah digunakan!"
            });
        }
        if (request.user_password !== request.user_confirmPassword) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Kata sandi tidak sesuai!"
            });
        }
    
        const hashedPassword = await bcrypt.hash(request.user_password, 10)
        const payload = { user_email: request.user_email, user_password: hashedPassword, user_created_at: new Date() }

        const user = new UserModels(payload)
        await user.save()
        return res.json({
            error: false,
            code: 200,
            message: 'Registrasi Berhasil',
            description: 'Berhasil melakukan registrasi!',
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

const Login = async (req, res) => {
    try {
        const { user_email, user_password } = req.body
        if (!user_email || !validator.isEmail(user_email) || !user_password) {
            return res.status(400).json({
                error: true,
                code: 400,
                message: 'Bad Request',
                description: "Permintaan tidak dapat diproses karena ada data yang hilang atau tidak valid!"
            });
        } 

        const user = await UserModels.findOne({ user_email })
        if (!user) {
            return res.status(401).json({ 
                error: true,
                code: 401,
                message: 'Invalid Request',
                description: 'Email tidak tersedia!',
            });
        }
        const passwordMatch = await bcrypt.compare(user_password, user.user_password);
        if (!passwordMatch) {
            return res.status(401).json({ 
                error: true,
                code: 401,
                message: 'Invalid Request',
                description: 'kata sandi tidak sesuai!',
            });
        }

        const key = process.env.API_KEY;
        const token = jwt.sign({ user_id: user._id, user_email: user.user_email, user_role: user.user_role }, key, { expiresIn: '3d' });

        return res.json({
            error: false,
            code: 200,
            message: 'Login Berhasil',
            description: 'Berhasil login!',
            data: token
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
    Register, Login
}