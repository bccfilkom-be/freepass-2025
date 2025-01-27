const express = require('express')
require('dotenv').config()
const { configDatabase } = require('./config');
const middleware = require('./middleware');
const bodyParser = require('body-parser');
const { Auth, Profile, User, Admin, Coordinator, Conference } = require('./routes');
const { ProfileController } = require('./controllers');

const App = express()
const PORT = process.env.PORT_API;

App.use(bodyParser.json());
App.use(bodyParser.urlencoded({ extended: false }));


App.get('/', (req, res)=>{
    res.json({
        message: "Hello World!",
        author: {
            nama: "Achmad Hasbil Wafi Rahmawan",
            prodi: "Sistem Informasi",
            angkatan: 2024,
            kontak: "https://wa.me/6289607810680"
        },
    })
})

App.use(middleware.apiKey)

App.use('/auth', Auth)
App.use('/user', middleware.authMiddleware('user'), User);
App.use('/admin', middleware.authMiddleware('admin'), Admin);
App.use('/coordinator', middleware.authMiddleware('coordinator'), Coordinator);
App.use('/profile', middleware.authMiddleware(['admin', 'user', "coordinator"]), Profile);
App.use('/search', middleware.authMiddleware(['admin', 'user', "coordinator"]), ProfileController.Search);
App.use('/conference', middleware.authMiddleware(['user', "coordinator"]), Conference);

App.listen(PORT, ()=>{
    console.log('====================================');
    console.log(`success running server in ${PORT}`)
    console.log('====================================');
    // connect database
    configDatabase()
})