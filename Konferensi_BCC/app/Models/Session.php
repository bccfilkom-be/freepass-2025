<?php

namespace App\Models;


use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;


class Session extends Model
{
    use HasFactory;

    
    protected $table = 'session';

    
    protected $fillable = [
        'title', 'description', 'author', 'start_time', 'end_time', 'capacity', 'participants'
    ];


    public function user()
    {
        return $this->belongsTo(User::class);
    }


    public function sessionRegist()
    {
        return $this->hasOne(SessionRegist::class);
    }


    public function feedback(){
        return $this->hasMany(Feedback::class);
    }
}
