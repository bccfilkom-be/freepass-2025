<?php

namespace App\Http\Controllers;


use App\Models\User;
use App\Models\Proposal;
use Illuminate\Http\Request;
use App\Models\SessionRegist;
use App\Http\Resources\UserResource;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;
use Illuminate\Routing\Controller;


class UserController extends Controller
{
    public function editUser( Request $request )
    {
        $user = Auth::user();


        $request->validate([
            'name' => 'required',  
            'email' => 'required',
            'password' => 'nullable', 
        ]);


        if( $request->role ) {
            return response()->json([
                'message' => 'You don\'t have permission to edit your role'
            ], 401);
        }


        $user->name = $request->name;
        $user->email = $request->email;


        if ( $request->password ) {
            $user->password = Hash::make($request->password);
        }


        $user->save();


        return response()->json([
            'message' => 'Profile updated successfully',
            'user' => new UserResource($user),
        ]);
    }


    public function searchUser( Request $request )
    {
        $existingUser = User::where('name', $request->name)->first();


        if( $existingUser ) {
            return response()->json([
                'user' => new UserResource($existingUser),
            ]);
        }


        else {
            return response()->json([
                'message' => 'User not found',
            ], 404);
        }
    }


    public function remove( Request $request, $id )
    {
        $user = User::find($id);


        if (!$user) {
            return response()->json(['message' => 'User not found.'], 404);
        }


        $proposal = Proposal::where('author', $id)->first();


    if ( $proposal ) {
        $proposal->delete();
    }


    $sessionRegist = SessionRegist::where('user_id', $id)->first();


    if ( $sessionRegist ) {
        $sessionRegist->delete();
    }


        $user->delete();


        return response()->json(['message' => 'User account deleted successfully.'], 200);
    }
}
