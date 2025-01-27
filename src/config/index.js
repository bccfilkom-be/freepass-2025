const mongoose = require('mongoose');
require('dotenv').config();

const uri = process.env.MONGO_URI;
const database = process.env.MONGO_DATABASE;

async function configDatabase() {
    try {
        await mongoose.connect(`${uri}/${database}`);
        console.log('====================================');
        console.log('Berhasil terhubung ke database');
        console.log('====================================');
    } catch (error) {
        console.error('====================================');
        console.error('Gagal terhubung ke database');
        console.error(error);
        console.error('====================================');
        throw error;
    }
}

module.exports = { configDatabase }
