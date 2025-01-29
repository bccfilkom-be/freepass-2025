<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;


return new class extends Migration
{
    public function up()
    {
        Schema::create('sessions', function (Blueprint $table) {
            $table->string('id')->primary();
            $table->string('user_id')->nullable();
            $table->string('ip_address')->nullable();
            $table->string('user_agent')->nullable(); 
            $table->text('payload');
            $table->integer('last_activity');
        });
    }


    public function down(): void
    {
        Schema::dropIfExists('sessions');
    }
};
