<?php

namespace App\Models;


use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Factories\HasFactory;


class Proposal extends Model
{
    use HasFactory;


    protected $table = 'proposals';

    
    protected $fillable = [
        'title', 'description', 'author', 'status'
    ];


    public function user()
    {
        return $this->belongsTo(User::class);
    }
}
