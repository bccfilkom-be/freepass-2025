<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;


return new class extends Migration
{
    public function up(): void
    {
        Schema::create('proposals', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('author');
            $table->string('title');
            $table->text('description');
            $table->enum('status', ['pending', 'accepted', 'rejected'])->default('pending');
            $table->timestamps();


            $table->foreign('author')->references('id')->on('users');
        });
    }

    
    public function down(): void
    {
        Schema::dropIfExists('proposals');
    }
};
