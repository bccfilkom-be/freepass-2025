<?php

namespace App\Http\Controllers;


use App\Http\Resources\FeedbackResource;
use App\Models\Feedback;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Routing\Controller;


class FeedbackController extends Controller
{
    public function create( Request $request, $session_id )
    {
        $request->validate([
            'feedback' => 'required'
        ]);


        $feedback = Feedback::create([
            'feedback' => $request->feedback,
            'session_id' => $session_id,
            'commenter' => Auth::id(),
        ]);


        return response()->json([
            'message' => 'Feedback created succesfully',
            'feedback' => new FeedbackResource($feedback)
        ], 201);
    }


    public function delete( $id )
    {
        $feedback = Feedback::findOrFail($id);
        $feedback->delete(); 


    return response()->json([
        'message' => "Feedback has been deleted successfully."
    ]);
    }
}
