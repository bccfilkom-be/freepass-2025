<?php

namespace App\Http\Middleware;


use Closure;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;


class AdminRole
{
    public function handle( Request $request, Closure $next ): Response
    {
        $currentUser = Auth::user();


        if( $currentUser->role === 'admin' ) {
            return $next( $request );
        }


        else {
            return response()->json(['message' => 'You don\t have permission to access this.'], 403);
        }
    }
}
