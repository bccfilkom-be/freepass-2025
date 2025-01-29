<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;


return new class extends Migration
{
    public function up()
    {
        Schema::table('session_registrations', function (Blueprint $table) {
            $table->timestamp('registration_at')->useCurrent()->change();
        });
    }


    public function down()
    {
        Schema::table('session_registrations', function (Blueprint $table) {
            $table->timestamp('registration_at')->nullable()->change();
        });
    }
};
