package main

// lmuVec3 represents a 3D vector.
type lmuVec3 struct {
	X float64
	Y float64
	Z float64
}

// lmuWheel represents a wheel of a vehicle.
type lmuWheel struct {
	SuspensionDeflection    float64
	RideHeight              float64
	SuspForce               float64
	BrakeTemp               float64
	BrakePressure           float64
	Rotation                float64
	LateralPatchVel         float64
	LongitudinalPatchVel    float64
	LateralGroundVel        float64
	LongitudinalGroundVel   float64
	Camber                  float64
	LateralForce            float64
	LongitudinalForce       float64
	TireLoad                float64
	GripFract               float64
	Pressure                float64
	Temperature             [3]float64
	Wear                    float64
	TerrainName             [16]byte
	SurfaceType             uint8
	Flat                    bool
	Detached                bool
	StaticUndeflectedRadius uint8
	VerticalTireDeflection  float64
	WheelYLocation          float64
	Toe                     float64
	TireCarcassTemperature  float64
	TireInnerLayerTemperature [3]float64
	Expansion               [24]byte
}

// lmuVehicleTelemetry represents telemetry data for a vehicle.
type lmuVehicleTelemetry struct {
	ID                       int32
	DeltaTime                float64
	ElapsedTime              float64
	LapNumber                int32
	LapStartET               float64
	VehicleName              [64]byte
	TrackName                [64]byte
	Pos                      lmuVec3
	LocalVel                 lmuVec3
	LocalAccel               lmuVec3
	Ori                      [3]lmuVec3
	LocalRot                 lmuVec3
	LocalRotAccel            lmuVec3
	Gear                     int32
	EngineRPM                float64
	EngineWaterTemp          float64
	EngineOilTemp            float64
	ClutchRPM                float64
	UnfilteredThrottle       float64
	UnfilteredBrake          float64
	UnfilteredSteering       float64
	UnfilteredClutch         float64
	FilteredThrottle         float64
	FilteredBrake            float64
	FilteredSteering         float64
	FilteredClutch           float64
	SteeringShaftTorque      float64
	Front3rdDeflection       float64
	Rear3rdDeflection        float64
	FrontWingHeight          float64
	FrontRideHeight          float64
	RearRideHeight           float64
	Drag                     float64
	FrontDownforce           float64
	RearDownforce            float64
	Fuel                     float64
	EngineMaxRPM             float64
	ScheduledStops           uint8
	Overheating              bool
	Detached                 bool
	Headlights               bool
	DentSeverity             [8]uint8
	LastImpactET             float64
	LastImpactMagnitude      float64
	LastImpactPos            lmuVec3
	EngineTorque             float64
	CurrentSector            int32
	SpeedLimiter             uint8
	MaxGears                 uint8
	FrontTireCompoundIndex   uint8
	RearTireCompoundIndex    uint8
	FuelCapacity             float64
	FrontFlapActivated       uint8
	RearFlapActivated        uint8
	RearFlapLegalStatus      uint8
	IgnitionStarter          uint8
	FrontTireCompoundName    [18]byte
	RearTireCompoundName     [18]byte
	SpeedLimiterAvailable    uint8
	AntiStallActivated       uint8
	Unused                   [2]uint8
	VisualSteeringWheelRange float32
	RearBrakeBias            float64
	TurboBoostPressure       float64
	PhysicsToGraphicsOffset  [3]float32
	PhysicalSteeringWheelRange float32
	DeltaBest                float64
	BatteryChargeFraction    float64
	ElectricBoostMotorTorque float64
	ElectricBoostMotorRPM    float64
	ElectricBoostMotorTemperature float64
	ElectricBoostWaterTemperature float64
	ElectricBoostMotorState  uint8
	Expansion                [103]byte
	Wheels                   [4]lmuWheel
}

// lmuScoringInfo represents scoring information.
type lmuScoringInfo struct {
	TrackName        [64]byte
	Session          int32
	CurrentET        float64
	EndET            float64
	MaxLaps          int32
	LapDist          float64
	Pointer1         [8]byte
	NumVehicles      int32
	GamePhase        uint8
	YellowFlagState  int8
	SectorFlag       [3]uint8
	StartLight       uint8
	NumRedLights     uint8
	InRealtime       bool
	PlayerName       [32]byte
	PlrFileName      [64]byte
	DarkCloud        float64
	Raining          float64
	AmbientTemp      float64
	TrackTemp        float64
	Wind             lmuVec3
	MinPathWetness   float64
	MaxPathWetness   float64
	GameMode         uint8
	IsPasswordProtected bool
	ServerPort       uint16
	ServerPublicIP   uint32
	MaxPlayers       int32
	ServerName       [32]byte
	StartET          float32
	AvgPathWetness   float64
	Expansion        [200]byte
	Pointer2         [8]byte
}

// lmuVehicleScoring represents scoring data for a vehicle.
type lmuVehicleScoring struct {
	ID                int32
	DriverName        [32]byte
	VehicleName       [64]byte
	TotalLaps         int16
	Sector            int8
	FinishStatus      int8
	LapDist           float64
	PathLateral       float64
	TrackEdge         float64
	BestSector1       float64
	BestSector2       float64
	BestLapTime       float64
	LastSector1       float64
	LastSector2       float64
	LastLapTime       float64
	CurSector1        float64
	CurSector2        float64
	NumPitstops       int16
	NumPenalties      int16
	IsPlayer          bool
	Control           int8
	InPits            bool
	Place             uint8
	VehicleClass      [32]byte
	TimeBehindNext    float64
	LapsBehindNext    int32
	TimeBehindLeader  float64
	LapsBehindLeader  int32
	LapStartET        float64
	Pos               lmuVec3
	LocalVel          lmuVec3
	LocalAccel        lmuVec3
	Ori               [3]lmuVec3
	LocalRot          lmuVec3
	LocalRotAccel     lmuVec3
	Headlights        uint8
	PitState          uint8
	ServerScored      uint8
	IndividualPhase   uint8
	Qualification     int32
	TimeIntoLap       float64
	EstimatedLapTime  float64
	PitGroup          [24]byte
	Flag              uint8
	UnderYellow       bool
	CountLapFlag      uint8
	InGarageStall     bool
	UpgradePack       [16]uint8
	PitLapDist        float32
	BestLapSector1    float32
	BestLapSector2    float32
	Expansion         [48]byte
}

// lmuTelemetry represents the entire telemetry data structure.
type lmuTelemetry struct {
	VersionUpdateBegin uint32
	VersionUpdateEnd   uint32
	BytesUpdatedHint   int32
	NumVehicles        int32
	Vehicles           [128]lmuVehicleTelemetry
}

// lmuScoring represents the entire scoring data structure.
type lmuScoring struct {
	VersionUpdateBegin uint32
	VersionUpdateEnd   uint32
	BytesUpdatedHint   int32
	ScoringInfo        lmuScoringInfo
	Vehicles           [128]lmuVehicleScoring
}
