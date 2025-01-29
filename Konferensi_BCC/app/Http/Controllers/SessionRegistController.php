<?php

namespace App\Http\Controllers;


use Carbon\Carbon;
use App\Models\Session;
use App\Models\SessionRegist;
use Illuminate\Support\Facades\Auth;
use App\Http\Resources\SessionResource;
use Illuminate\Routing\Controller;


class SessionRegistController extends Controller
{
    public function index( $session_id )
    {
        $session = Session::findOrFail($session_id);

        if ( $session->participants >= $session->capacity ) {
            return response()->json([
                'message' => 'Failed to register for the session because it is full.',
                'session' => new SessionResource($session)
            ], 400);
        }


        $session->increment('participants');


        $session_regist = SessionRegist::create([
            'session_id' => $session_id,
            'user_id' => Auth::id(),
            'registration_at' => Carbon::now(),
        ]);


        return response()->json([
            'message' => 'Successfully registered for the session.',
            'session' => new SessionResource($session),
            'registration_at' => $session_regist->registration_at->format('Y-m-d H:i:s')
        ]);
    }
}
