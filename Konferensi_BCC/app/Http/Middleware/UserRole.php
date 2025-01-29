<?php

namespace App\Http\Middleware;


use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;


class UserRole
{
    public function handle( Request $request, Closure $next ): Response
    {
        $currentUser = Auth::user();


        if( $currentUser->role === 'admin' ){
            return $next( $request );
        }


        else {
            if( $request->id != $currentUser->id ) {
                return response()->json(['message' => 'You don\'t have permission to delete this account'], 403);
            }

            
            return $next( $request );
        }
    }
}
