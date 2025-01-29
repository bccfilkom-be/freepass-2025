<?php

use GuzzleHttp\Middleware;
use Illuminate\Http\Request;
use App\Http\Middleware\UserRole;
use App\Http\Middleware\AdminRole;
use App\Http\Middleware\SessionRole;
use App\Http\Middleware\FeedbackRole;
use App\Http\Middleware\ProposalRole;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UserController;
use App\Http\Controllers\AdminController;
use App\Http\Controllers\SessionController;
use App\Http\Controllers\FeedbackController;
use App\Http\Controllers\ProposalController;
use App\Http\Controllers\SessionRegistController;
use App\Http\Controllers\AuthenticationController;


Route::post('/register', [AuthenticationController::class, 'register']);
Route::post('/login', [AuthenticationController::class, 'login']);
Route::get('/logout', [AuthenticationController::class, 'logout'])->middleware(['auth:sanctum']);


Route::put('/edit-user', [UserController::class, 'editUser'])->middleware('auth:sanctum');
Route::get('/search-user', [UserController::class, 'searchUser'])->middleware('auth:sanctum');
Route::delete('/remove/{id}', [UserController::class, 'remove'])->middleware(['auth:sanctum', UserRole::class]);


Route::post('/add-event-coordinator', [AdminController::class, 'createEventCoordinator'])->middleware(['auth:sanctum', AdminRole::class]);


Route::get('session', [SessionController::class, 'index'])->middleware(['auth:sanctum']);
Route::get('session/{id}', [SessionController::class, 'detail'])->middleware(['auth:sanctum', SessionRole::class]);
Route::put('session/{id}', [SessionController::class, 'update'])->middleware(['auth:sanctum', SessionRole::class]);
Route::delete('session/{id}', [SessionController::class, 'delete'])->middleware(['auth:sanctum', SessionRole::class]);


Route::get('/proposal', [ProposalController::class, 'index'])->middleware('auth:sanctum', ProposalRole::class);
Route::get('/proposal/{id}', [ProposalController::class, 'detail'])->middleware('auth:sanctum', ProposalRole::class);
Route::post('/proposal', [ProposalController::class, 'create'])->middleware('auth:sanctum', ProposalRole::class);
Route::put('/proposal/{id}', [ProposalController::class, 'update'])->middleware('auth:sanctum', ProposalRole::class);
Route::delete('proposal/{id}', [ProposalController::class, 'delete'])->middleware('auth:sanctum', ProposalRole::class);


Route::post('/feedback/{session_id}', [FeedbackController::class, 'create'])->middleware('auth:sanctum', FeedbackRole::class);
Route::delete('feedback/{id}', [FeedbackController::class, 'delete'])->middleware(['auth:sanctum', FeedbackRole::class]);


Route::get('/session-registration/{id}', [SessionRegistController::class, 'index'])->middleware('auth:sanctum');
