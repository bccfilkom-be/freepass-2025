<?php

namespace App\Http\Middleware;


use Closure;
use App\Models\Session;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;


class SessionRole
{
    public function handle( Request $request, Closure $next )
    {
        $currentUser = Auth::user();


        if( $currentUser->role === 'user' ) {
            $session = Session::findOrFail($request->id);
    

            if( !$session || $session->author != $currentUser->id ) {
                return response()->json(['message' => 'Data not found'], 404);
            }
    

            return $next( $request );
        }


        else if( $currentUser->role === 'event_coordinator') {
            if( $request->isMethod('put') ) {
                return response()->json(['message' => 'You don\'t have a permission to update this session'], 403);
            }
            
            
            else {
                return $next( $request );
            }
        }


        else {
            return response()->json(['message' => 'You don\'t have a permission to update this session'], 403);
        }
    }
}
