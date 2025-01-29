<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;


return new class extends Migration
{
    public function up()
    {
        Schema::table('sessions', function (Blueprint $table) {
            $table->dropForeign(['author']); 
        });


        Schema::table('session_registrations', function (Blueprint $table) {
            $table->dropForeign(['session_id']);
        });


        Schema::table('feedbacks', function (Blueprint $table) {
            $table->dropForeign(['session_id']);
        });


        Schema::rename('sessions', 'session');


        Schema::table('session', function (Blueprint $table) {
            $table->foreign('author')->references('id')->on('users');
        });


        Schema::table('session_registrations', function (Blueprint $table) {
            $table->foreign('session_id')->references('id')->on('session');
        });


        Schema::table('feedbacks', function (Blueprint $table) {
            $table->foreign('session_id')->references('id')->on('session');
        });
    }


    public function down()
    {
        Schema::rename('session', 'sessions');


        Schema::table('session_registrations', function (Blueprint $table) {
            $table->dropForeign(['session_id']);
        });


        Schema::table('feedbacks', function (Blueprint $table) {
            $table->dropForeign(['session_id']);
        });


        Schema::table('sessions', function (Blueprint $table) {
            $table->dropForeign(['author']);
        });


        Schema::table('sessions', function (Blueprint $table) {
            $table->foreign('author')->references('id')->on('users');
        });


        Schema::table('session_registrations', function (Blueprint $table) {
            $table->foreign('session_id')->references('id')->on('sessions');
        });

        
        Schema::table('feedbacks', function (Blueprint $table) {
            $table->foreign('session_id')->references('id')->on('sessions');
        });
    }
};
