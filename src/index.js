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
        API: {
            headers: {
                "x-api-key": "freepast-2025"
            },
            Authorization: "Barrier Token"
        },
        endPoint: {
            register: {
                method: "post",
                url: "/register",
                params: ["user_email", "user_password", "user_confirmPassword"]
            },
            login: {
                method: "post",
                url: "/login",
                params: ["user_email", "user_password"]
            },
            updateProfile: {
                method: "post",
                url: "/login",
                params: ["user_nama", "user_username"]
            },
            searchProfile: {
                method: "get",
                url: "/login",
                params: ["user_id", "user_username"]
            }
        },
        system: [
            "New users can register accounts in the system ✔️ (done)",
            "Users can log in to the system ✔️ (done)",
            "Users can edit their profile accounts ✔️ (done)",
            "Users can view all conference sessions ✔️ (done)",
            "Users can leave feedback on sessions ✔️ (done)",
            "Users can view other users' profiles ✔️ (done)",
            "Users can register for sessions during the conference registration period if seats are available ✔️ (done)",
            "Users can only register for one session within a time period ✔️ (done)",
            "Users can create, edit, and delete their session proposals ✔️ (done)",
            "Users can only create one session proposal within a time period ✔️ (done)",
            "Users can edit and delete their sessions ✔️ (done)",
            "Event Coordinators can view all session proposals ✔️ (done)",
            "Event Coordinators can accept or reject user session proposals ✔️ (done)",
            "Event Coordinators can remove sessions ✔️ (done)",
            "Event Coordinators can remove user feedback ✔️ (done)",
            "Admins can add new event coordinators ✔️ (done)",
            "Admins can remove users or event coordinators ✔️ (done)",
        ]
        
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