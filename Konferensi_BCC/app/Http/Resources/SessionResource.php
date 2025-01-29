<?php

namespace App\Http\Resources;


use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;


class SessionResource extends JsonResource
{
    public function toArray( Request $request ): array
    {
        return [
            'title' => $this->title,
            'description' => $this->description,
            'author' => $this->author,
            'start_time' => $this->start_time,
            'end_time' => $this->end_time,
            'capacity' => $this->capacity,
            'feedbacks' => FeedbackResource::collection($this->feedback)
        ];
    }
}
