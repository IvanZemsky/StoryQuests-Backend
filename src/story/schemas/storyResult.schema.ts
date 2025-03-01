import { Prop, Schema, SchemaFactory } from "@nestjs/mongoose"
import { User } from "src/user/user.schema"
import { Story } from "./story.schema"
import mongoose, { HydratedDocument } from "mongoose"

@Schema({
   collection: "story-results",
   versionKey: false,
})
export class StoryResult {
   _id: string

   @Prop({ required: true })
   resultSceneNumber: string

   @Prop({ required: true, ref: "User", type: mongoose.Schema.Types.ObjectId })
   userId: User

   @Prop({ required: true, ref: "Story", type: mongoose.Schema.Types.ObjectId })
   storyId: Story

   @Prop({ required: true })
   datetime: Date
}

export type StoryResultDocument = HydratedDocument<StoryResult>
export const StoryResultSchema = SchemaFactory.createForClass(StoryResult)
