package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type User struct {
	Id             types.Int64    `tfsdk:"id,omitempty"`
	Status         types.Int64    `tfsdk:"status"`
	Username       types.String   `tfsdk:"username"`
	Description    types.String   `tfsdk:"description"`
	ExpirationDate types.Float64  `tfsdk:"expiration_date"`
	Password       types.String   `tfsdk:"password,omitempty"`
	PublicKeys     []types.String `tfsdk:"public_keys,omitempty"`
	HomeDir        types.String   `tfsdk:"home_dir"`
	//Permissions       map[types.String][]types.String `tfsdk:"permissions"`
	Uid               types.Int64     `tfsdk:"uid"`
	Gid               types.Int64     `tfsdk:"gid"`
	MaxSessions       types.Int64     `tfsdk:"max_sessions"`
	QuotaSize         types.Float64   `tfsdk:"quota_size"`
	QuotaFiles        types.Int64     `tfsdk:"quota_files"`
	VirtualFolders    []VirtualFolder `tfsdk:"virtual_folders"`
	UploadBandwidth   types.Int64     `tfsdk:"upload_bandwidth"`
	DownloadBandwidth types.Int64     `tfsdk:"download_bandwidth"`
	Filters           *Filters        `tfsdk:"filters,omitempty"`
	Filesystem        *Filesystem     `tfsdk:"filesystem,omitempty"`
	AdditionalInfo    types.String    `tfsdk:"additional_info"`
	LastUpdated       types.String    `tfsdk:"last_updated"`
}

type VirtualFolder struct {
	Id          types.Int64    `tfsdk:"id"`
	Name        types.String   `tfsdk:"name"`
	MappedPath  types.String   `tfsdk:"mapped_path"`
	Users       []types.String `tfsdk:"users"`
	Description types.String   `tfsdk:"description"`
	QuotaSize   types.Int64    `tfsdk:"quota_size"`
	QuotaFiles  types.Int64    `tfsdk:"quota_files"`
	VirtualPath types.String   `tfsdk:"virtual_path"`
	// Just for mapping. Not used
	UsedQuotaSize types.Int64 `tfsdk:"used_quota_size"`
	// Just for mapping. Not used
	UsedQuotaFiles types.Int64 `tfsdk:"used_quota_files"`
	// Just for mapping. Not used
	LastQuotaUpdate types.Int64 `tfsdk:"last_quota_update"`
	Filesystem      *Filesystem `tfsdk:"filesystem,omitempty"`
}

type Filesystem struct {
	Provider     types.Int64   `tfsdk:"provider"`
	S3Config     *S3Config     `tfsdk:"s3config,omitempty"`
	Gcsconfig    *Gcsconfig    `tfsdk:"gcsconfig,omitempty"`
	Azblobconfig *Azblobconfig `tfsdk:"azblobconfig,omitempty"`
	Cryptconfig  *Cryptconfig  `tfsdk:"cryptconfig,omitempty"`
	Sftpconfig   *Sftpconfig   `tfsdk:"sftpconfig,omitempty"`
}

type SftpPassword struct {
	Status         types.String `tfsdk:"status"`
	Payload        types.String `tfsdk:"payload"`
	Key            types.String `tfsdk:"key"`
	AdditionalData types.String `tfsdk:"additional_data"`
	Mode           types.Int64  `tfsdk:"mode"`
}

type SftpPrivateKey struct {
	Status         types.String `tfsdk:"status"`
	Payload        types.String `tfsdk:"payload"`
	Key            types.String `tfsdk:"key"`
	AdditionalData types.String `tfsdk:"additional_data"`
	Mode           types.Int64  `tfsdk:"mode"`
}

type Sftpconfig struct {
	Endpotypes             types.String   `tfsdk:"endpotypes.Int64"`
	Username               types.String   `tfsdk:"username"`
	Password               SftpPassword   `tfsdk:"password"`
	PrivateKey             SftpPrivateKey `tfsdk:"private_key"`
	Fingerprtypes          []types.String `tfsdk:"fingerprtypes.Int64s"`
	Prefix                 types.String   `tfsdk:"prefix"`
	DisableConcurrentReads bool           `tfsdk:"disable_concurrent_reads"`
	BufferSize             types.Int64    `tfsdk:"buffer_size"`
}

type CryptPassphrase struct {
	Status         types.String `tfsdk:"status"`
	Payload        types.String `tfsdk:"payload"`
	Key            types.String `tfsdk:"key"`
	AdditionalData types.String `tfsdk:"additional_data"`
	Mode           types.Int64  `tfsdk:"mode"`
}

type Cryptconfig struct {
	Passphrase CryptPassphrase `tfsdk:"passphrase"`
}

type AzAccountKey struct {
	Status         types.String `tfsdk:"status"`
	Payload        types.String `tfsdk:"payload"`
	Key            types.String `tfsdk:"key"`
	AdditionalData types.String `tfsdk:"additional_data"`
	Mode           types.Int64  `tfsdk:"mode"`
}

type Azblobconfig struct {
	Container         types.String `tfsdk:"container"`
	AccountName       types.String `tfsdk:"account_name"`
	AccountKey        AzAccountKey `tfsdk:"account_key"`
	SasUrl            types.String `tfsdk:"sas_url"`
	Endpotypes        types.String `tfsdk:"endpotypes.Int64"`
	UploadPartSize    types.Int64  `tfsdk:"upload_part_size"`
	UploadConcurrency types.Int64  `tfsdk:"upload_concurrency"`
	AccessTier        types.String `tfsdk:"access_tier"`
	KeyPrefix         types.String `tfsdk:"key_prefix"`
	UseEmulator       bool         `tfsdk:"use_emulator"`
}

type GcsCredentials struct {
	Status         types.String `tfsdk:"status"`
	Payload        types.String `tfsdk:"payload"`
	Key            types.String `tfsdk:"key"`
	AdditionalData types.String `tfsdk:"additional_data"`
	Mode           types.Int64  `tfsdk:"mode"`
}

type Gcsconfig struct {
	Bucket               types.String   `tfsdk:"bucket"`
	Credentials          GcsCredentials `tfsdk:"credentials,omitempty"`
	AutomaticCredentials types.Int64    `tfsdk:"automatic_credentials"`
	StorageClass         types.String   `tfsdk:"storage_class"`
	KeyPrefix            types.String   `tfsdk:"key_prefix"`
}

type S3AccessSecret struct {
	Status         types.String `tfsdk:"status"`
	Payload        types.String `tfsdk:"payload"`
	Key            types.String `tfsdk:"key"`
	AdditionalData types.String `tfsdk:"additional_data"`
	Mode           types.Int64  `tfsdk:"mode"`
}

type S3Config struct {
	Bucket            types.String   `tfsdk:"bucket"`
	Region            types.String   `tfsdk:"region"`
	AccessKey         types.String   `tfsdk:"access_key"`
	AccessSecret      S3AccessSecret `tfsdk:"access_secret"`
	Endpotypes        types.String   `tfsdk:"endpotypes.Int64"`
	StorageClass      types.String   `tfsdk:"storage_class"`
	UploadPartSize    types.Int64    `tfsdk:"upload_part_size"`
	UploadConcurrency types.Int64    `tfsdk:"upload_concurrency"`
	KeyPrefix         types.String   `tfsdk:"key_prefix"`
}

type FilePatterns struct {
	Path            types.String   `tfsdk:"path"`
	AllowedPatterns []types.String `tfsdk:"allowed_patterns"`
	DeniedPatterns  []types.String `tfsdk:"denied_patterns"`
}

type Filters struct {
	AllowedIp          []types.String `tfsdk:"allowed_ip"`
	DeniedIp           []types.String `tfsdk:"denied_ip"`
	DeniedLoginMethods []types.String `tfsdk:"denied_login_methods"`
	DeniedProtocols    []types.String `tfsdk:"denied_protocols"`
	FilePatterns       []FilePatterns `tfsdk:"file_patterns"`
}
