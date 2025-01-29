<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;


return new class extends Migration
{
    public function up(): void
    {
        Schema::create('feedbacks', function (Blueprint $table) {
            $table->id();
            $table->unsignedBigInteger('commenter');
            $table->unsignedBigInteger('session_id');
            $table->text('feedback');
            $table->timestamps();


            $table->foreign('commenter')->references('id')->on('users');
            $table->foreign('session_id')->references('id')->on('sessions');
        });
    }

    public function down(): void
    {
        Schema::dropIfExists('feedbacks');
    }
};
