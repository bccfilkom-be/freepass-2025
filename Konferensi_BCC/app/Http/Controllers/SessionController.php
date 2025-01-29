<?php

namespace App\Http\Controllers;


use App\Models\Session;
use App\Models\Feedback;
use Illuminate\Http\Request;
use App\Models\SessionRegist;
use Illuminate\Routing\Controller;
use App\Http\Resources\SessionResource;


class SessionController extends Controller
{
    public function index()
    {
        $sessions = Session::all();
        return SessionResource::collection($sessions);
    }


    public function detail( $id )
    {
        $session = Session::findOrFail($id);
        return response()->json([
            'session' => new SessionResource($session)
        ]); 
    }
    

    public function update( Request $request, $id )
    {
        $session = Session::findOrFail($id);
        $session->update( $request->all() );
        return response()->json([
            'session' => new SessionResource($session)
        ]);     
    }
    

    public function delete( $id )
    {
        $feedbacks = Feedback::where('session_id', $id)->get();
        if ($feedbacks->count() > 0) {
            $feedbacks->each->delete();
        }


        $sessionRegistrations = SessionRegist::where('session_id', $id)->get();
        if ($sessionRegistrations->count() > 0) {
            $sessionRegistrations->each->delete();
        }


        $session = Session::findOrFail($id);
        $title = $session->title; 
        $session->delete(); 


        return response()->json([
            'message' => "Session with title '{$title}' has been deleted successfully."
        ]);
    }
}
