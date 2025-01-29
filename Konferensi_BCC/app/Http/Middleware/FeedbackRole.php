<?php

namespace App\Http\Middleware;


use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;


class FeedbackRole
{
    public function handle( Request $request, Closure $next ): Response
    {
        $currentUser = Auth::user();


        if( $currentUser->role === 'event_coordinator' ) {
            if( $request->isMethod('delete') ) {
                return $next( $request );
            }


            else {
                return response()->json(['message' => 'Your don\'t have permission to create feedbacks'], 403);
            }
        }


        else if( $currentUser->role === 'user' ) {
            if( $request->isMethod('post') ) {
                return $next( $request );
            }


            else {
                if( $currentUser != $request->commenter ) {
                    return response()->json(['message' => 'Feedback has deleted succesfully'
                    ]);
                }
                
                
                else {
                    return response()->json(['message' => 'Data not found'], 404);
                }
            }
        }
        
        
        else {
            return response()->json(['message' => 'You don\'t have permission to create or delete feedbacks'], 403);
        }
    }
}
