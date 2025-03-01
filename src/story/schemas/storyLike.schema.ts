import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose"
import { ApiProperty } from "@nestjs/swagger"
import mongoose, { HydratedDocument, ObjectId } from "mongoose"

@Schema({ collection: "story-likes", versionKey: false })
export class StoryLike {
   @ApiProperty({
      example: "66cb6fb8ebae2e4b8fffd190",
      description: "Уникальный идентификатор",
   })
   _id: ObjectId

   @ApiProperty({
      example: "66cb6fb8ebae2e4b8fffd190",
      description: "Уникальный идентификатор истории",
   })
   @Prop({ type: mongoose.Schema.Types.ObjectId })
   storyId: ObjectId

   @ApiProperty({
      example: "66cb6fb8ebae2e4b8fffd190",
      description: "Уникальный идентификатор пользователя",
   })
   @Prop({ type: mongoose.Schema.Types.ObjectId })
   userId: ObjectId
}

export type StoryLikeDocument = HydratedDocument<StoryLike>
export const StoryLikeSchema = SchemaFactory.createForClass(StoryLike)
