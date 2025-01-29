<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;


return new class extends Migration
{
    public function up(): void
    {
        Schema::table('session', function (Blueprint $table) {
            $table->integer('participants')->default(0);
        });
    }


    public function down(): void
    {
        Schema::table('session', function (Blueprint $table) {
            $table->dropColumn('participants');
        });
    }
};
