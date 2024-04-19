package models

import (
  "time"

  "github.com/google/uuid"
)

type Player struct {
  ID                       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
  AccelerationRating       uint8     `json:"acceleration_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  Age                      uint8     `json:"age" binding:"numeric,min=0,max=99" gorm:"not null"`
  AgilityRating            uint8     `json:"agility_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  AwarenessRating          uint8     `json:"awareness_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  BallCarrierVisionRating  uint8     `json:"ball_carrier_vision_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  BirthDate                time.Time `json:"birth_date" binding:"required,max=30" gorm:"type:date;not null"`
  BlockSheddingRating      uint8     `json:"block_shedding_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  BreakSackRating          uint8     `json:"break_sack_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  BreakTackleRating        uint8     `json:"break_tackle_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  CarryingRating           uint8     `json:"carrying_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  CatchInTrafficRating     uint8     `json:"catch_in_traffic_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  CatchingRating           uint8     `json:"catching_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ChangeOfDirectionRating  uint8     `json:"change_of_direction_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  College                  string    `json:"college" binding:"required,max=100" gorm:"type:varchar(100);not null"`
  DeepRouteRunningRating   uint8     `json:"deep_route_running_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  FinesseMovesRating       uint8     `json:"finesse_moves_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  FirstName                string    `json:"first_name" binding:"required,max=30" gorm:"type:varchar(30);not null"`
  Handedness               string    `json:"handedness" binding:"required,max=10" gorm:"type:varchar(10);not null"`
  Height                   uint8     `json:"height" binding:"numeric,min=1,max=99" gorm:"not null"`
  HitPowerRating           uint8     `json:"hit_power_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ImpactBlockingRating     uint8     `json:"impact_blocking_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  InjuryRating             uint8     `json:"injury_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  JerseyNum                uint8     `json:"jersey_num" binding:"numeric,min=0,max=99" gorm:"not null"`
  JukeMoveRating           uint8     `json:"juke_move_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  JumpingRating            uint8     `json:"jumping_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  KickAccuracyRating       uint8     `json:"kick_accuracy_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  KickPowerRating          uint8     `json:"kick_power_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  KickReturnRating         uint8     `json:"kick_return_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  LastName                 string    `json:"last_name" binding:"required,max=30" gorm:"type:varchar(30);not null"`
  LeadBlockRating          uint8     `json:"lead_block_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ManCoverageRating        uint8     `json:"man_coverage_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  MediumRouteRunningRating uint8     `json:"medium_route_running_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  OverallRating            uint8     `json:"overall_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PassBlockFinesseRating   uint8     `json:"pass_block_finesse_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PassBlockPowerRating     uint8     `json:"pass_block_power_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PassBlockRating          uint8     `json:"pass_block_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PlayActionRating         uint8     `json:"play_action_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PlayRecognitionRating    uint8     `json:"play_recognition_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  Position                 string    `json:"position" binding:"required,max=5" gorm:"type:varchar(5);not null"`
  PowerMovesRating         uint8     `json:"power_moves_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PressRating              uint8     `json:"press_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  PursuitRating            uint8     `json:"pursuit_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ReleaseRating            uint8     `json:"release_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  RunBlockFinesseRating    uint8     `json:"run_block_finesse_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  RunBlockPowerRating      uint8     `json:"run_block_power_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  RunBlockRating           uint8     `json:"run_block_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  RunningStyleRating       string    `json:"running_style_rating" binding:"required,max=40" gorm:"type:varchar(40);not null"`
  ShortRouteRunningRating  uint8     `json:"short_route_running_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  SpectacularCatchRating   uint8     `json:"spectacular_catch_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  SpeedRating              uint8     `json:"speed_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  SpinMoveRating           uint8     `json:"spin_move_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  StaminaRating            uint8     `json:"stamina_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  StiffArmRating           uint8     `json:"stiff_arm_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  StrengthRating           uint8     `json:"strength_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  TackleRating             uint8     `json:"tackle_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  Team                     string    `json:"team" binding:"required,max=20" gorm:"type:varchar(20);not null"`
  ThrowAccuracyDeepRating  uint8     `json:"throw_accuracy_deep_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ThrowAccuracyMidRating   uint8     `json:"throw_accuracy_mid_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ThrowAccuracyShortRating uint8     `json:"throw_accuracy_short_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ThrowOnTheRunRating      uint8     `json:"throw_on_the_run_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ThrowPowerRating         uint8     `json:"throw_power_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ThrowUnderPressureRating uint8     `json:"throw_under_pressure_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  ToughnessRating          uint8     `json:"toughness_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  TruckingRating           uint8     `json:"trucking_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
  Weight                   uint16    `json:"weight" binding:"numeric,min=1,max=500" gorm:"not null"`
  NumYears                 uint8     `json:"num_years" binding:"numeric,min=0,max=99" gorm:"not null"`
  ZoneCoverageRating       uint8     `json:"zone_coverage_rating" binding:"numeric,min=0,max=99" gorm:"not null"`
}
