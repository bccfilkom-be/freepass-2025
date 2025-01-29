<?php

namespace App\Http\Controllers;


use App\Models\User;
use Illuminate\Http\Request;
use App\Http\Resources\UserResource;
use Illuminate\Support\Facades\Hash;
use Illuminate\Routing\Controller;


class AuthenticationController extends Controller
{
    public function register( Request $request )
    {
        $request->validate([
            'name' => 'required',
            'email' => 'required|email',
            'password' => 'required',
            'role' => 'required|in:user,admin,event_coordinator'
        ]);


        $existingUser = User::where('email', $request->email)->first();


        if ( $existingUser ) {
            return response()->json([
                'message' => 'The email is already taken.',
                'user' => new UserResource($existingUser),
            ], 409);
        }


        $user = User::create([
            'name' => $request->name,
            'email' => $request->email,
            'password' => Hash::make($request->password),
            'role' => $request->role, 
        ]);


        return response()->json([
            'message' => 'User registered successfully',
            'user' => new UserResource($user),
        ], 201);
    }


    public function login( Request $request )
    {
        $request->validate([
            'email' => 'required|email',
            'password' => 'required',
            'role' => 'required|in:user,admin,event_coordinator'
        ]);


        $user = User::where('email', $request->email)->first();


        if ( !$user || ! Hash::check($request->password, $user->password )) {
            return response()->json(['message' => 'Invalid credentials.'], 401);
        }


        $token = $user->createToken('API Token')->plainTextToken;


        return response()->json([
            'message' => 'Login successful',
            'user' => new UserResource($user),
            'token' => $token,
        ]);
    }

    
    public function logout( Request $request )
    {
        $user = $request->user();
        $user->currentAccessToken()->delete();


        return response()->json([
            'message' => 'Logout success',
            'user' => new UserResource($user),
        ]);
    }
}
