<?php

namespace App\Http\Middleware;


use Closure;
use App\Models\Proposal;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Symfony\Component\HttpFoundation\Response;


class ProposalRole
{
    public function handle( Request $request, Closure $next ): Response
    {
        $currentUser = Auth::user();


        if ( $currentUser->role === 'user' ) {
                if ( $request->id ) {
                    $proposal = Proposal::find($request->id);
    
    
                    if ( !$proposal || $proposal->author != $currentUser->id ) {
                        return response()->json(['message' => 'Data not found'], 404);
                    }
                }

                
                return $next( $request );
        }


        else if ( $currentUser->role === 'event_coordinator' ) {
            if ( $request->isMethod('get') || $request->isMethod('put') ) {
                return $next( $request );
            }


            return response()->json(['message' => 'You don\'t have permission to create or delete proposals'], 403);
        }

        else {
            return response()->json(['message' => 'Unauthorized'], 403);
        }
    }
}
