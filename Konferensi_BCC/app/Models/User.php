<?php

namespace App\Models;


use Laravel\Sanctum\HasApiTokens;
use Illuminate\Notifications\Notifiable;
use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;


class User extends Authenticatable
{
    use HasFactory, Notifiable, HasApiTokens;


    protected $fillable = [
        'name',
        'email',
        'password',
        'role'
    ];

    
    protected $hidden = [
        'password',
        'remember_token',
    ];

    
    protected function casts(): array
    {
        return [
            'email_verified_at' => 'datetime',
            'password' => 'hashed',
        ];
    }


    public function session(){
        return $this->hasMany(Session::class);
    }


    public function sessionRegist()
    {
        return $this->hasOne(SessionRegist::class);
    }


    public function proposal()
    {
        return $this->hasMany(Proposal::class);
    }


    public function feedback()
    {
        return $this->hasMany(Feedback::class);
    }
}
