<?php

namespace App\Models;


use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;


class SessionRegist extends Model
{
    use HasFactory;


    protected $table = 'session_registrations';

    
    protected $fillable = [
        'user_id', 'session_id', 'registration_at'
    ];


    public function user(){
        return $this->hasOne(User::class);
    }


    public function session()
    {
        return $this->hasOne(Session::class);
    }
}
