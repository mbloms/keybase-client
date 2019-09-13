// Auto-generated to Go types and interfaces using avdl-compiler v1.4.2 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/simple_fs.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type OpID [16]byte

func (o OpID) DeepCopy() OpID {
	var ret OpID
	copy(ret[:], o[:])
	return ret
}

type KBFSRevision int64

func (o KBFSRevision) DeepCopy() KBFSRevision {
	return o
}

type KBFSArchivedType int

const (
	KBFSArchivedType_REVISION        KBFSArchivedType = 0
	KBFSArchivedType_TIME            KBFSArchivedType = 1
	KBFSArchivedType_TIME_STRING     KBFSArchivedType = 2
	KBFSArchivedType_REL_TIME_STRING KBFSArchivedType = 3
)

func (o KBFSArchivedType) DeepCopy() KBFSArchivedType { return o }

var KBFSArchivedTypeMap = map[string]KBFSArchivedType{
	"REVISION":        0,
	"TIME":            1,
	"TIME_STRING":     2,
	"REL_TIME_STRING": 3,
}

var KBFSArchivedTypeRevMap = map[KBFSArchivedType]string{
	0: "REVISION",
	1: "TIME",
	2: "TIME_STRING",
	3: "REL_TIME_STRING",
}

func (e KBFSArchivedType) String() string {
	if v, ok := KBFSArchivedTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type KBFSArchivedParam struct {
	KBFSArchivedType__ KBFSArchivedType `codec:"KBFSArchivedType" json:"KBFSArchivedType"`
	Revision__         *KBFSRevision    `codec:"revision,omitempty" json:"revision,omitempty"`
	Time__             *Time            `codec:"time,omitempty" json:"time,omitempty"`
	TimeString__       *string          `codec:"timeString,omitempty" json:"timeString,omitempty"`
	RelTimeString__    *string          `codec:"relTimeString,omitempty" json:"relTimeString,omitempty"`
}

func (o *KBFSArchivedParam) KBFSArchivedType() (ret KBFSArchivedType, err error) {
	switch o.KBFSArchivedType__ {
	case KBFSArchivedType_REVISION:
		if o.Revision__ == nil {
			err = errors.New("unexpected nil value for Revision__")
			return ret, err
		}
	case KBFSArchivedType_TIME:
		if o.Time__ == nil {
			err = errors.New("unexpected nil value for Time__")
			return ret, err
		}
	case KBFSArchivedType_TIME_STRING:
		if o.TimeString__ == nil {
			err = errors.New("unexpected nil value for TimeString__")
			return ret, err
		}
	case KBFSArchivedType_REL_TIME_STRING:
		if o.RelTimeString__ == nil {
			err = errors.New("unexpected nil value for RelTimeString__")
			return ret, err
		}
	}
	return o.KBFSArchivedType__, nil
}

func (o KBFSArchivedParam) Revision() (res KBFSRevision) {
	if o.KBFSArchivedType__ != KBFSArchivedType_REVISION {
		panic("wrong case accessed")
	}
	if o.Revision__ == nil {
		return
	}
	return *o.Revision__
}

func (o KBFSArchivedParam) Time() (res Time) {
	if o.KBFSArchivedType__ != KBFSArchivedType_TIME {
		panic("wrong case accessed")
	}
	if o.Time__ == nil {
		return
	}
	return *o.Time__
}

func (o KBFSArchivedParam) TimeString() (res string) {
	if o.KBFSArchivedType__ != KBFSArchivedType_TIME_STRING {
		panic("wrong case accessed")
	}
	if o.TimeString__ == nil {
		return
	}
	return *o.TimeString__
}

func (o KBFSArchivedParam) RelTimeString() (res string) {
	if o.KBFSArchivedType__ != KBFSArchivedType_REL_TIME_STRING {
		panic("wrong case accessed")
	}
	if o.RelTimeString__ == nil {
		return
	}
	return *o.RelTimeString__
}

func NewKBFSArchivedParamWithRevision(v KBFSRevision) KBFSArchivedParam {
	return KBFSArchivedParam{
		KBFSArchivedType__: KBFSArchivedType_REVISION,
		Revision__:         &v,
	}
}

func NewKBFSArchivedParamWithTime(v Time) KBFSArchivedParam {
	return KBFSArchivedParam{
		KBFSArchivedType__: KBFSArchivedType_TIME,
		Time__:             &v,
	}
}

func NewKBFSArchivedParamWithTimeString(v string) KBFSArchivedParam {
	return KBFSArchivedParam{
		KBFSArchivedType__: KBFSArchivedType_TIME_STRING,
		TimeString__:       &v,
	}
}

func NewKBFSArchivedParamWithRelTimeString(v string) KBFSArchivedParam {
	return KBFSArchivedParam{
		KBFSArchivedType__: KBFSArchivedType_REL_TIME_STRING,
		RelTimeString__:    &v,
	}
}

func (o KBFSArchivedParam) DeepCopy() KBFSArchivedParam {
	return KBFSArchivedParam{
		KBFSArchivedType__: o.KBFSArchivedType__.DeepCopy(),
		Revision__: (func(x *KBFSRevision) *KBFSRevision {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Revision__),
		Time__: (func(x *Time) *Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Time__),
		TimeString__: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.TimeString__),
		RelTimeString__: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.RelTimeString__),
	}
}

type KBFSArchivedPath struct {
	Path             string               `codec:"path" json:"path"`
	ArchivedParam    KBFSArchivedParam    `codec:"archivedParam" json:"archivedParam"`
	IdentifyBehavior *TLFIdentifyBehavior `codec:"identifyBehavior,omitempty" json:"identifyBehavior,omitempty"`
}

func (o KBFSArchivedPath) DeepCopy() KBFSArchivedPath {
	return KBFSArchivedPath{
		Path:          o.Path,
		ArchivedParam: o.ArchivedParam.DeepCopy(),
		IdentifyBehavior: (func(x *TLFIdentifyBehavior) *TLFIdentifyBehavior {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.IdentifyBehavior),
	}
}

type KBFSPath struct {
	Path             string               `codec:"path" json:"path"`
	IdentifyBehavior *TLFIdentifyBehavior `codec:"identifyBehavior,omitempty" json:"identifyBehavior,omitempty"`
}

func (o KBFSPath) DeepCopy() KBFSPath {
	return KBFSPath{
		Path: o.Path,
		IdentifyBehavior: (func(x *TLFIdentifyBehavior) *TLFIdentifyBehavior {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.IdentifyBehavior),
	}
}

type PathType int

const (
	PathType_LOCAL         PathType = 0
	PathType_KBFS          PathType = 1
	PathType_KBFS_ARCHIVED PathType = 2
)

func (o PathType) DeepCopy() PathType { return o }

var PathTypeMap = map[string]PathType{
	"LOCAL":         0,
	"KBFS":          1,
	"KBFS_ARCHIVED": 2,
}

var PathTypeRevMap = map[PathType]string{
	0: "LOCAL",
	1: "KBFS",
	2: "KBFS_ARCHIVED",
}

func (e PathType) String() string {
	if v, ok := PathTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type Path struct {
	PathType__     PathType          `codec:"PathType" json:"PathType"`
	Local__        *string           `codec:"local,omitempty" json:"local,omitempty"`
	Kbfs__         *KBFSPath         `codec:"kbfs,omitempty" json:"kbfs,omitempty"`
	KbfsArchived__ *KBFSArchivedPath `codec:"kbfsArchived,omitempty" json:"kbfsArchived,omitempty"`
}

func (o *Path) PathType() (ret PathType, err error) {
	switch o.PathType__ {
	case PathType_LOCAL:
		if o.Local__ == nil {
			err = errors.New("unexpected nil value for Local__")
			return ret, err
		}
	case PathType_KBFS:
		if o.Kbfs__ == nil {
			err = errors.New("unexpected nil value for Kbfs__")
			return ret, err
		}
	case PathType_KBFS_ARCHIVED:
		if o.KbfsArchived__ == nil {
			err = errors.New("unexpected nil value for KbfsArchived__")
			return ret, err
		}
	}
	return o.PathType__, nil
}

func (o Path) Local() (res string) {
	if o.PathType__ != PathType_LOCAL {
		panic("wrong case accessed")
	}
	if o.Local__ == nil {
		return
	}
	return *o.Local__
}

func (o Path) Kbfs() (res KBFSPath) {
	if o.PathType__ != PathType_KBFS {
		panic("wrong case accessed")
	}
	if o.Kbfs__ == nil {
		return
	}
	return *o.Kbfs__
}

func (o Path) KbfsArchived() (res KBFSArchivedPath) {
	if o.PathType__ != PathType_KBFS_ARCHIVED {
		panic("wrong case accessed")
	}
	if o.KbfsArchived__ == nil {
		return
	}
	return *o.KbfsArchived__
}

func NewPathWithLocal(v string) Path {
	return Path{
		PathType__: PathType_LOCAL,
		Local__:    &v,
	}
}

func NewPathWithKbfs(v KBFSPath) Path {
	return Path{
		PathType__: PathType_KBFS,
		Kbfs__:     &v,
	}
}

func NewPathWithKbfsArchived(v KBFSArchivedPath) Path {
	return Path{
		PathType__:     PathType_KBFS_ARCHIVED,
		KbfsArchived__: &v,
	}
}

func (o Path) DeepCopy() Path {
	return Path{
		PathType__: o.PathType__.DeepCopy(),
		Local__: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Local__),
		Kbfs__: (func(x *KBFSPath) *KBFSPath {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Kbfs__),
		KbfsArchived__: (func(x *KBFSArchivedPath) *KBFSArchivedPath {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.KbfsArchived__),
	}
}

type DirentType int

const (
	DirentType_FILE DirentType = 0
	DirentType_DIR  DirentType = 1
	DirentType_SYM  DirentType = 2
	DirentType_EXEC DirentType = 3
)

func (o DirentType) DeepCopy() DirentType { return o }

var DirentTypeMap = map[string]DirentType{
	"FILE": 0,
	"DIR":  1,
	"SYM":  2,
	"EXEC": 3,
}

var DirentTypeRevMap = map[DirentType]string{
	0: "FILE",
	1: "DIR",
	2: "SYM",
	3: "EXEC",
}

func (e DirentType) String() string {
	if v, ok := DirentTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type PrefetchStatus int

const (
	PrefetchStatus_NOT_STARTED PrefetchStatus = 0
	PrefetchStatus_IN_PROGRESS PrefetchStatus = 1
	PrefetchStatus_COMPLETE    PrefetchStatus = 2
)

func (o PrefetchStatus) DeepCopy() PrefetchStatus { return o }

var PrefetchStatusMap = map[string]PrefetchStatus{
	"NOT_STARTED": 0,
	"IN_PROGRESS": 1,
	"COMPLETE":    2,
}

var PrefetchStatusRevMap = map[PrefetchStatus]string{
	0: "NOT_STARTED",
	1: "IN_PROGRESS",
	2: "COMPLETE",
}

func (e PrefetchStatus) String() string {
	if v, ok := PrefetchStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type PrefetchProgress struct {
	Start        Time  `codec:"start" json:"start"`
	EndEstimate  Time  `codec:"endEstimate" json:"endEstimate"`
	BytesTotal   int64 `codec:"bytesTotal" json:"bytesTotal"`
	BytesFetched int64 `codec:"bytesFetched" json:"bytesFetched"`
}

func (o PrefetchProgress) DeepCopy() PrefetchProgress {
	return PrefetchProgress{
		Start:        o.Start.DeepCopy(),
		EndEstimate:  o.EndEstimate.DeepCopy(),
		BytesTotal:   o.BytesTotal,
		BytesFetched: o.BytesFetched,
	}
}

type Dirent struct {
	Time                 Time             `codec:"time" json:"time"`
	Size                 int              `codec:"size" json:"size"`
	Name                 string           `codec:"name" json:"name"`
	DirentType           DirentType       `codec:"direntType" json:"direntType"`
	LastWriterUnverified User             `codec:"lastWriterUnverified" json:"lastWriterUnverified"`
	Writable             bool             `codec:"writable" json:"writable"`
	PrefetchStatus       PrefetchStatus   `codec:"prefetchStatus" json:"prefetchStatus"`
	PrefetchProgress     PrefetchProgress `codec:"prefetchProgress" json:"prefetchProgress"`
}

func (o Dirent) DeepCopy() Dirent {
	return Dirent{
		Time:                 o.Time.DeepCopy(),
		Size:                 o.Size,
		Name:                 o.Name,
		DirentType:           o.DirentType.DeepCopy(),
		LastWriterUnverified: o.LastWriterUnverified.DeepCopy(),
		Writable:             o.Writable,
		PrefetchStatus:       o.PrefetchStatus.DeepCopy(),
		PrefetchProgress:     o.PrefetchProgress.DeepCopy(),
	}
}

type DirentWithRevision struct {
	Entry    Dirent       `codec:"entry" json:"entry"`
	Revision KBFSRevision `codec:"revision" json:"revision"`
}

func (o DirentWithRevision) DeepCopy() DirentWithRevision {
	return DirentWithRevision{
		Entry:    o.Entry.DeepCopy(),
		Revision: o.Revision.DeepCopy(),
	}
}

type RevisionSpanType int

const (
	RevisionSpanType_DEFAULT   RevisionSpanType = 0
	RevisionSpanType_LAST_FIVE RevisionSpanType = 1
)

func (o RevisionSpanType) DeepCopy() RevisionSpanType { return o }

var RevisionSpanTypeMap = map[string]RevisionSpanType{
	"DEFAULT":   0,
	"LAST_FIVE": 1,
}

var RevisionSpanTypeRevMap = map[RevisionSpanType]string{
	0: "DEFAULT",
	1: "LAST_FIVE",
}

func (e RevisionSpanType) String() string {
	if v, ok := RevisionSpanTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type ErrorNum int

func (o ErrorNum) DeepCopy() ErrorNum {
	return o
}

type OpenFlags int

const (
	OpenFlags_READ      OpenFlags = 0
	OpenFlags_REPLACE   OpenFlags = 1
	OpenFlags_EXISTING  OpenFlags = 2
	OpenFlags_WRITE     OpenFlags = 4
	OpenFlags_APPEND    OpenFlags = 8
	OpenFlags_DIRECTORY OpenFlags = 16
)

func (o OpenFlags) DeepCopy() OpenFlags { return o }

var OpenFlagsMap = map[string]OpenFlags{
	"READ":      0,
	"REPLACE":   1,
	"EXISTING":  2,
	"WRITE":     4,
	"APPEND":    8,
	"DIRECTORY": 16,
}

var OpenFlagsRevMap = map[OpenFlags]string{
	0:  "READ",
	1:  "REPLACE",
	2:  "EXISTING",
	4:  "WRITE",
	8:  "APPEND",
	16: "DIRECTORY",
}

func (e OpenFlags) String() string {
	if v, ok := OpenFlagsRevMap[e]; ok {
		return v
	}
	return ""
}

type Progress int

func (o Progress) DeepCopy() Progress {
	return o
}

type SimpleFSListResult struct {
	Entries  []Dirent `codec:"entries" json:"entries"`
	Progress Progress `codec:"progress" json:"progress"`
}

func (o SimpleFSListResult) DeepCopy() SimpleFSListResult {
	return SimpleFSListResult{
		Entries: (func(x []Dirent) []Dirent {
			if x == nil {
				return nil
			}
			ret := make([]Dirent, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Entries),
		Progress: o.Progress.DeepCopy(),
	}
}

type FileContent struct {
	Data     []byte   `codec:"data" json:"data"`
	Progress Progress `codec:"progress" json:"progress"`
}

func (o FileContent) DeepCopy() FileContent {
	return FileContent{
		Data: (func(x []byte) []byte {
			if x == nil {
				return nil
			}
			return append([]byte{}, x...)
		})(o.Data),
		Progress: o.Progress.DeepCopy(),
	}
}

type AsyncOps int

const (
	AsyncOps_LIST                    AsyncOps = 0
	AsyncOps_LIST_RECURSIVE          AsyncOps = 1
	AsyncOps_READ                    AsyncOps = 2
	AsyncOps_WRITE                   AsyncOps = 3
	AsyncOps_COPY                    AsyncOps = 4
	AsyncOps_MOVE                    AsyncOps = 5
	AsyncOps_REMOVE                  AsyncOps = 6
	AsyncOps_LIST_RECURSIVE_TO_DEPTH AsyncOps = 7
	AsyncOps_GET_REVISIONS           AsyncOps = 8
)

func (o AsyncOps) DeepCopy() AsyncOps { return o }

var AsyncOpsMap = map[string]AsyncOps{
	"LIST":                    0,
	"LIST_RECURSIVE":          1,
	"READ":                    2,
	"WRITE":                   3,
	"COPY":                    4,
	"MOVE":                    5,
	"REMOVE":                  6,
	"LIST_RECURSIVE_TO_DEPTH": 7,
	"GET_REVISIONS":           8,
}

var AsyncOpsRevMap = map[AsyncOps]string{
	0: "LIST",
	1: "LIST_RECURSIVE",
	2: "READ",
	3: "WRITE",
	4: "COPY",
	5: "MOVE",
	6: "REMOVE",
	7: "LIST_RECURSIVE_TO_DEPTH",
	8: "GET_REVISIONS",
}

func (e AsyncOps) String() string {
	if v, ok := AsyncOpsRevMap[e]; ok {
		return v
	}
	return ""
}

type ListFilter int

const (
	ListFilter_NO_FILTER            ListFilter = 0
	ListFilter_FILTER_ALL_HIDDEN    ListFilter = 1
	ListFilter_FILTER_SYSTEM_HIDDEN ListFilter = 2
)

func (o ListFilter) DeepCopy() ListFilter { return o }

var ListFilterMap = map[string]ListFilter{
	"NO_FILTER":            0,
	"FILTER_ALL_HIDDEN":    1,
	"FILTER_SYSTEM_HIDDEN": 2,
}

var ListFilterRevMap = map[ListFilter]string{
	0: "NO_FILTER",
	1: "FILTER_ALL_HIDDEN",
	2: "FILTER_SYSTEM_HIDDEN",
}

func (e ListFilter) String() string {
	if v, ok := ListFilterRevMap[e]; ok {
		return v
	}
	return ""
}

type ListArgs struct {
	OpID   OpID       `codec:"opID" json:"opID"`
	Path   Path       `codec:"path" json:"path"`
	Filter ListFilter `codec:"filter" json:"filter"`
}

func (o ListArgs) DeepCopy() ListArgs {
	return ListArgs{
		OpID:   o.OpID.DeepCopy(),
		Path:   o.Path.DeepCopy(),
		Filter: o.Filter.DeepCopy(),
	}
}

type ListToDepthArgs struct {
	OpID   OpID       `codec:"opID" json:"opID"`
	Path   Path       `codec:"path" json:"path"`
	Filter ListFilter `codec:"filter" json:"filter"`
	Depth  int        `codec:"depth" json:"depth"`
}

func (o ListToDepthArgs) DeepCopy() ListToDepthArgs {
	return ListToDepthArgs{
		OpID:   o.OpID.DeepCopy(),
		Path:   o.Path.DeepCopy(),
		Filter: o.Filter.DeepCopy(),
		Depth:  o.Depth,
	}
}

type RemoveArgs struct {
	OpID      OpID `codec:"opID" json:"opID"`
	Path      Path `codec:"path" json:"path"`
	Recursive bool `codec:"recursive" json:"recursive"`
}

func (o RemoveArgs) DeepCopy() RemoveArgs {
	return RemoveArgs{
		OpID:      o.OpID.DeepCopy(),
		Path:      o.Path.DeepCopy(),
		Recursive: o.Recursive,
	}
}

type ReadArgs struct {
	OpID   OpID  `codec:"opID" json:"opID"`
	Path   Path  `codec:"path" json:"path"`
	Offset int64 `codec:"offset" json:"offset"`
	Size   int   `codec:"size" json:"size"`
}

func (o ReadArgs) DeepCopy() ReadArgs {
	return ReadArgs{
		OpID:   o.OpID.DeepCopy(),
		Path:   o.Path.DeepCopy(),
		Offset: o.Offset,
		Size:   o.Size,
	}
}

type WriteArgs struct {
	OpID   OpID  `codec:"opID" json:"opID"`
	Path   Path  `codec:"path" json:"path"`
	Offset int64 `codec:"offset" json:"offset"`
}

func (o WriteArgs) DeepCopy() WriteArgs {
	return WriteArgs{
		OpID:   o.OpID.DeepCopy(),
		Path:   o.Path.DeepCopy(),
		Offset: o.Offset,
	}
}

type CopyArgs struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

func (o CopyArgs) DeepCopy() CopyArgs {
	return CopyArgs{
		OpID: o.OpID.DeepCopy(),
		Src:  o.Src.DeepCopy(),
		Dest: o.Dest.DeepCopy(),
	}
}

type MoveArgs struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

func (o MoveArgs) DeepCopy() MoveArgs {
	return MoveArgs{
		OpID: o.OpID.DeepCopy(),
		Src:  o.Src.DeepCopy(),
		Dest: o.Dest.DeepCopy(),
	}
}

type GetRevisionsArgs struct {
	OpID     OpID             `codec:"opID" json:"opID"`
	Path     Path             `codec:"path" json:"path"`
	SpanType RevisionSpanType `codec:"spanType" json:"spanType"`
}

func (o GetRevisionsArgs) DeepCopy() GetRevisionsArgs {
	return GetRevisionsArgs{
		OpID:     o.OpID.DeepCopy(),
		Path:     o.Path.DeepCopy(),
		SpanType: o.SpanType.DeepCopy(),
	}
}

type OpDescription struct {
	AsyncOp__              AsyncOps          `codec:"asyncOp" json:"asyncOp"`
	List__                 *ListArgs         `codec:"list,omitempty" json:"list,omitempty"`
	ListRecursive__        *ListArgs         `codec:"listRecursive,omitempty" json:"listRecursive,omitempty"`
	ListRecursiveToDepth__ *ListToDepthArgs  `codec:"listRecursiveToDepth,omitempty" json:"listRecursiveToDepth,omitempty"`
	Read__                 *ReadArgs         `codec:"read,omitempty" json:"read,omitempty"`
	Write__                *WriteArgs        `codec:"write,omitempty" json:"write,omitempty"`
	Copy__                 *CopyArgs         `codec:"copy,omitempty" json:"copy,omitempty"`
	Move__                 *MoveArgs         `codec:"move,omitempty" json:"move,omitempty"`
	Remove__               *RemoveArgs       `codec:"remove,omitempty" json:"remove,omitempty"`
	GetRevisions__         *GetRevisionsArgs `codec:"getRevisions,omitempty" json:"getRevisions,omitempty"`
}

func (o *OpDescription) AsyncOp() (ret AsyncOps, err error) {
	switch o.AsyncOp__ {
	case AsyncOps_LIST:
		if o.List__ == nil {
			err = errors.New("unexpected nil value for List__")
			return ret, err
		}
	case AsyncOps_LIST_RECURSIVE:
		if o.ListRecursive__ == nil {
			err = errors.New("unexpected nil value for ListRecursive__")
			return ret, err
		}
	case AsyncOps_LIST_RECURSIVE_TO_DEPTH:
		if o.ListRecursiveToDepth__ == nil {
			err = errors.New("unexpected nil value for ListRecursiveToDepth__")
			return ret, err
		}
	case AsyncOps_READ:
		if o.Read__ == nil {
			err = errors.New("unexpected nil value for Read__")
			return ret, err
		}
	case AsyncOps_WRITE:
		if o.Write__ == nil {
			err = errors.New("unexpected nil value for Write__")
			return ret, err
		}
	case AsyncOps_COPY:
		if o.Copy__ == nil {
			err = errors.New("unexpected nil value for Copy__")
			return ret, err
		}
	case AsyncOps_MOVE:
		if o.Move__ == nil {
			err = errors.New("unexpected nil value for Move__")
			return ret, err
		}
	case AsyncOps_REMOVE:
		if o.Remove__ == nil {
			err = errors.New("unexpected nil value for Remove__")
			return ret, err
		}
	case AsyncOps_GET_REVISIONS:
		if o.GetRevisions__ == nil {
			err = errors.New("unexpected nil value for GetRevisions__")
			return ret, err
		}
	}
	return o.AsyncOp__, nil
}

func (o OpDescription) List() (res ListArgs) {
	if o.AsyncOp__ != AsyncOps_LIST {
		panic("wrong case accessed")
	}
	if o.List__ == nil {
		return
	}
	return *o.List__
}

func (o OpDescription) ListRecursive() (res ListArgs) {
	if o.AsyncOp__ != AsyncOps_LIST_RECURSIVE {
		panic("wrong case accessed")
	}
	if o.ListRecursive__ == nil {
		return
	}
	return *o.ListRecursive__
}

func (o OpDescription) ListRecursiveToDepth() (res ListToDepthArgs) {
	if o.AsyncOp__ != AsyncOps_LIST_RECURSIVE_TO_DEPTH {
		panic("wrong case accessed")
	}
	if o.ListRecursiveToDepth__ == nil {
		return
	}
	return *o.ListRecursiveToDepth__
}

func (o OpDescription) Read() (res ReadArgs) {
	if o.AsyncOp__ != AsyncOps_READ {
		panic("wrong case accessed")
	}
	if o.Read__ == nil {
		return
	}
	return *o.Read__
}

func (o OpDescription) Write() (res WriteArgs) {
	if o.AsyncOp__ != AsyncOps_WRITE {
		panic("wrong case accessed")
	}
	if o.Write__ == nil {
		return
	}
	return *o.Write__
}

func (o OpDescription) Copy() (res CopyArgs) {
	if o.AsyncOp__ != AsyncOps_COPY {
		panic("wrong case accessed")
	}
	if o.Copy__ == nil {
		return
	}
	return *o.Copy__
}

func (o OpDescription) Move() (res MoveArgs) {
	if o.AsyncOp__ != AsyncOps_MOVE {
		panic("wrong case accessed")
	}
	if o.Move__ == nil {
		return
	}
	return *o.Move__
}

func (o OpDescription) Remove() (res RemoveArgs) {
	if o.AsyncOp__ != AsyncOps_REMOVE {
		panic("wrong case accessed")
	}
	if o.Remove__ == nil {
		return
	}
	return *o.Remove__
}

func (o OpDescription) GetRevisions() (res GetRevisionsArgs) {
	if o.AsyncOp__ != AsyncOps_GET_REVISIONS {
		panic("wrong case accessed")
	}
	if o.GetRevisions__ == nil {
		return
	}
	return *o.GetRevisions__
}

func NewOpDescriptionWithList(v ListArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_LIST,
		List__:    &v,
	}
}

func NewOpDescriptionWithListRecursive(v ListArgs) OpDescription {
	return OpDescription{
		AsyncOp__:       AsyncOps_LIST_RECURSIVE,
		ListRecursive__: &v,
	}
}

func NewOpDescriptionWithListRecursiveToDepth(v ListToDepthArgs) OpDescription {
	return OpDescription{
		AsyncOp__:              AsyncOps_LIST_RECURSIVE_TO_DEPTH,
		ListRecursiveToDepth__: &v,
	}
}

func NewOpDescriptionWithRead(v ReadArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_READ,
		Read__:    &v,
	}
}

func NewOpDescriptionWithWrite(v WriteArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_WRITE,
		Write__:   &v,
	}
}

func NewOpDescriptionWithCopy(v CopyArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_COPY,
		Copy__:    &v,
	}
}

func NewOpDescriptionWithMove(v MoveArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_MOVE,
		Move__:    &v,
	}
}

func NewOpDescriptionWithRemove(v RemoveArgs) OpDescription {
	return OpDescription{
		AsyncOp__: AsyncOps_REMOVE,
		Remove__:  &v,
	}
}

func NewOpDescriptionWithGetRevisions(v GetRevisionsArgs) OpDescription {
	return OpDescription{
		AsyncOp__:      AsyncOps_GET_REVISIONS,
		GetRevisions__: &v,
	}
}

func (o OpDescription) DeepCopy() OpDescription {
	return OpDescription{
		AsyncOp__: o.AsyncOp__.DeepCopy(),
		List__: (func(x *ListArgs) *ListArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.List__),
		ListRecursive__: (func(x *ListArgs) *ListArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ListRecursive__),
		ListRecursiveToDepth__: (func(x *ListToDepthArgs) *ListToDepthArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ListRecursiveToDepth__),
		Read__: (func(x *ReadArgs) *ReadArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Read__),
		Write__: (func(x *WriteArgs) *WriteArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Write__),
		Copy__: (func(x *CopyArgs) *CopyArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Copy__),
		Move__: (func(x *MoveArgs) *MoveArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Move__),
		Remove__: (func(x *RemoveArgs) *RemoveArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Remove__),
		GetRevisions__: (func(x *GetRevisionsArgs) *GetRevisionsArgs {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.GetRevisions__),
	}
}

type GetRevisionsResult struct {
	Revisions []DirentWithRevision `codec:"revisions" json:"revisions"`
	Progress  Progress             `codec:"progress" json:"progress"`
}

func (o GetRevisionsResult) DeepCopy() GetRevisionsResult {
	return GetRevisionsResult{
		Revisions: (func(x []DirentWithRevision) []DirentWithRevision {
			if x == nil {
				return nil
			}
			ret := make([]DirentWithRevision, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Revisions),
		Progress: o.Progress.DeepCopy(),
	}
}

type OpProgress struct {
	Start        Time     `codec:"start" json:"start"`
	EndEstimate  Time     `codec:"endEstimate" json:"endEstimate"`
	OpType       AsyncOps `codec:"opType" json:"opType"`
	BytesTotal   int64    `codec:"bytesTotal" json:"bytesTotal"`
	BytesRead    int64    `codec:"bytesRead" json:"bytesRead"`
	BytesWritten int64    `codec:"bytesWritten" json:"bytesWritten"`
	FilesTotal   int64    `codec:"filesTotal" json:"filesTotal"`
	FilesRead    int64    `codec:"filesRead" json:"filesRead"`
	FilesWritten int64    `codec:"filesWritten" json:"filesWritten"`
}

func (o OpProgress) DeepCopy() OpProgress {
	return OpProgress{
		Start:        o.Start.DeepCopy(),
		EndEstimate:  o.EndEstimate.DeepCopy(),
		OpType:       o.OpType.DeepCopy(),
		BytesTotal:   o.BytesTotal,
		BytesRead:    o.BytesRead,
		BytesWritten: o.BytesWritten,
		FilesTotal:   o.FilesTotal,
		FilesRead:    o.FilesRead,
		FilesWritten: o.FilesWritten,
	}
}

type SimpleFSGetHTTPAddressAndTokenResponse struct {
	Address string `codec:"address" json:"address"`
	Token   string `codec:"token" json:"token"`
}

func (o SimpleFSGetHTTPAddressAndTokenResponse) DeepCopy() SimpleFSGetHTTPAddressAndTokenResponse {
	return SimpleFSGetHTTPAddressAndTokenResponse{
		Address: o.Address,
		Token:   o.Token,
	}
}

type SimpleFSQuotaUsage struct {
	UsageBytes      int64 `codec:"usageBytes" json:"usageBytes"`
	ArchiveBytes    int64 `codec:"archiveBytes" json:"archiveBytes"`
	LimitBytes      int64 `codec:"limitBytes" json:"limitBytes"`
	GitUsageBytes   int64 `codec:"gitUsageBytes" json:"gitUsageBytes"`
	GitArchiveBytes int64 `codec:"gitArchiveBytes" json:"gitArchiveBytes"`
	GitLimitBytes   int64 `codec:"gitLimitBytes" json:"gitLimitBytes"`
}

func (o SimpleFSQuotaUsage) DeepCopy() SimpleFSQuotaUsage {
	return SimpleFSQuotaUsage{
		UsageBytes:      o.UsageBytes,
		ArchiveBytes:    o.ArchiveBytes,
		LimitBytes:      o.LimitBytes,
		GitUsageBytes:   o.GitUsageBytes,
		GitArchiveBytes: o.GitArchiveBytes,
		GitLimitBytes:   o.GitLimitBytes,
	}
}

type FolderSyncMode int

const (
	FolderSyncMode_DISABLED FolderSyncMode = 0
	FolderSyncMode_ENABLED  FolderSyncMode = 1
	FolderSyncMode_PARTIAL  FolderSyncMode = 2
)

func (o FolderSyncMode) DeepCopy() FolderSyncMode { return o }

var FolderSyncModeMap = map[string]FolderSyncMode{
	"DISABLED": 0,
	"ENABLED":  1,
	"PARTIAL":  2,
}

var FolderSyncModeRevMap = map[FolderSyncMode]string{
	0: "DISABLED",
	1: "ENABLED",
	2: "PARTIAL",
}

func (e FolderSyncMode) String() string {
	if v, ok := FolderSyncModeRevMap[e]; ok {
		return v
	}
	return ""
}

type FolderSyncConfig struct {
	Mode  FolderSyncMode `codec:"mode" json:"mode"`
	Paths []string       `codec:"paths" json:"paths"`
}

func (o FolderSyncConfig) DeepCopy() FolderSyncConfig {
	return FolderSyncConfig{
		Mode: o.Mode.DeepCopy(),
		Paths: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Paths),
	}
}

type FolderSyncConfigAndStatus struct {
	Config FolderSyncConfig `codec:"config" json:"config"`
	Status FolderSyncStatus `codec:"status" json:"status"`
}

func (o FolderSyncConfigAndStatus) DeepCopy() FolderSyncConfigAndStatus {
	return FolderSyncConfigAndStatus{
		Config: o.Config.DeepCopy(),
		Status: o.Status.DeepCopy(),
	}
}

type FolderSyncConfigAndStatusWithFolder struct {
	Folder Folder           `codec:"folder" json:"folder"`
	Config FolderSyncConfig `codec:"config" json:"config"`
	Status FolderSyncStatus `codec:"status" json:"status"`
}

func (o FolderSyncConfigAndStatusWithFolder) DeepCopy() FolderSyncConfigAndStatusWithFolder {
	return FolderSyncConfigAndStatusWithFolder{
		Folder: o.Folder.DeepCopy(),
		Config: o.Config.DeepCopy(),
		Status: o.Status.DeepCopy(),
	}
}

type SyncConfigAndStatusRes struct {
	Folders       []FolderSyncConfigAndStatusWithFolder `codec:"folders" json:"folders"`
	OverallStatus FolderSyncStatus                      `codec:"overallStatus" json:"overallStatus"`
}

func (o SyncConfigAndStatusRes) DeepCopy() SyncConfigAndStatusRes {
	return SyncConfigAndStatusRes{
		Folders: (func(x []FolderSyncConfigAndStatusWithFolder) []FolderSyncConfigAndStatusWithFolder {
			if x == nil {
				return nil
			}
			ret := make([]FolderSyncConfigAndStatusWithFolder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Folders),
		OverallStatus: o.OverallStatus.DeepCopy(),
	}
}

type FSSettings struct {
	SpaceAvailableNotificationThreshold int64 `codec:"spaceAvailableNotificationThreshold" json:"spaceAvailableNotificationThreshold"`
}

func (o FSSettings) DeepCopy() FSSettings {
	return FSSettings{
		SpaceAvailableNotificationThreshold: o.SpaceAvailableNotificationThreshold,
	}
}

type SimpleFSStats struct {
	ProcessStats      ProcessRuntimeStats `codec:"processStats" json:"processStats"`
	BlockCacheDbStats []string            `codec:"blockCacheDbStats" json:"blockCacheDbStats"`
	SyncCacheDbStats  []string            `codec:"syncCacheDbStats" json:"syncCacheDbStats"`
	RuntimeDbStats    []DbStats           `codec:"runtimeDbStats" json:"runtimeDbStats"`
}

func (o SimpleFSStats) DeepCopy() SimpleFSStats {
	return SimpleFSStats{
		ProcessStats: o.ProcessStats.DeepCopy(),
		BlockCacheDbStats: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.BlockCacheDbStats),
		SyncCacheDbStats: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.SyncCacheDbStats),
		RuntimeDbStats: (func(x []DbStats) []DbStats {
			if x == nil {
				return nil
			}
			ret := make([]DbStats, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.RuntimeDbStats),
	}
}

type SubscriptionTopic int

const (
	SubscriptionTopic_FAVORITES      SubscriptionTopic = 0
	SubscriptionTopic_JOURNAL_STATUS SubscriptionTopic = 1
	SubscriptionTopic_ONLINE_STATUS  SubscriptionTopic = 2
)

func (o SubscriptionTopic) DeepCopy() SubscriptionTopic { return o }

var SubscriptionTopicMap = map[string]SubscriptionTopic{
	"FAVORITES":      0,
	"JOURNAL_STATUS": 1,
	"ONLINE_STATUS":  2,
}

var SubscriptionTopicRevMap = map[SubscriptionTopic]string{
	0: "FAVORITES",
	1: "JOURNAL_STATUS",
	2: "ONLINE_STATUS",
}

func (e SubscriptionTopic) String() string {
	if v, ok := SubscriptionTopicRevMap[e]; ok {
		return v
	}
	return ""
}

type PathSubscriptionTopic int

const (
	PathSubscriptionTopic_CHILDREN PathSubscriptionTopic = 0
	PathSubscriptionTopic_STAT     PathSubscriptionTopic = 1
)

func (o PathSubscriptionTopic) DeepCopy() PathSubscriptionTopic { return o }

var PathSubscriptionTopicMap = map[string]PathSubscriptionTopic{
	"CHILDREN": 0,
	"STAT":     1,
}

var PathSubscriptionTopicRevMap = map[PathSubscriptionTopic]string{
	0: "CHILDREN",
	1: "STAT",
}

func (e PathSubscriptionTopic) String() string {
	if v, ok := PathSubscriptionTopicRevMap[e]; ok {
		return v
	}
	return ""
}

type SimpleFSListArg struct {
	OpID                OpID       `codec:"opID" json:"opID"`
	Path                Path       `codec:"path" json:"path"`
	Filter              ListFilter `codec:"filter" json:"filter"`
	RefreshSubscription bool       `codec:"refreshSubscription" json:"refreshSubscription"`
}

type SimpleFSListRecursiveArg struct {
	OpID                OpID       `codec:"opID" json:"opID"`
	Path                Path       `codec:"path" json:"path"`
	Filter              ListFilter `codec:"filter" json:"filter"`
	RefreshSubscription bool       `codec:"refreshSubscription" json:"refreshSubscription"`
}

type SimpleFSListRecursiveToDepthArg struct {
	OpID                OpID       `codec:"opID" json:"opID"`
	Path                Path       `codec:"path" json:"path"`
	Filter              ListFilter `codec:"filter" json:"filter"`
	RefreshSubscription bool       `codec:"refreshSubscription" json:"refreshSubscription"`
	Depth               int        `codec:"depth" json:"depth"`
}

type SimpleFSReadListArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSCopyArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSSymlinkArg struct {
	Target string `codec:"target" json:"target"`
	Link   Path   `codec:"link" json:"link"`
}

type SimpleFSCopyRecursiveArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSMoveArg struct {
	OpID OpID `codec:"opID" json:"opID"`
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSRenameArg struct {
	Src  Path `codec:"src" json:"src"`
	Dest Path `codec:"dest" json:"dest"`
}

type SimpleFSOpenArg struct {
	OpID  OpID      `codec:"opID" json:"opID"`
	Dest  Path      `codec:"dest" json:"dest"`
	Flags OpenFlags `codec:"flags" json:"flags"`
}

type SimpleFSSetStatArg struct {
	Dest Path       `codec:"dest" json:"dest"`
	Flag DirentType `codec:"flag" json:"flag"`
}

type SimpleFSReadArg struct {
	OpID   OpID  `codec:"opID" json:"opID"`
	Offset int64 `codec:"offset" json:"offset"`
	Size   int   `codec:"size" json:"size"`
}

type SimpleFSWriteArg struct {
	OpID    OpID   `codec:"opID" json:"opID"`
	Offset  int64  `codec:"offset" json:"offset"`
	Content []byte `codec:"content" json:"content"`
}

type SimpleFSRemoveArg struct {
	OpID      OpID `codec:"opID" json:"opID"`
	Path      Path `codec:"path" json:"path"`
	Recursive bool `codec:"recursive" json:"recursive"`
}

type SimpleFSStatArg struct {
	Path                Path `codec:"path" json:"path"`
	RefreshSubscription bool `codec:"refreshSubscription" json:"refreshSubscription"`
}

type SimpleFSGetRevisionsArg struct {
	OpID     OpID             `codec:"opID" json:"opID"`
	Path     Path             `codec:"path" json:"path"`
	SpanType RevisionSpanType `codec:"spanType" json:"spanType"`
}

type SimpleFSReadRevisionsArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSMakeOpidArg struct {
}

type SimpleFSCloseArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSCancelArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSCheckArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSGetOpsArg struct {
}

type SimpleFSWaitArg struct {
	OpID OpID `codec:"opID" json:"opID"`
}

type SimpleFSDumpDebuggingInfoArg struct {
}

type SimpleFSClearConflictStateArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSFinishResolvingConflictArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSForceStuckConflictArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSSyncStatusArg struct {
	Filter ListFilter `codec:"filter" json:"filter"`
}

type SimpleFSGetHTTPAddressAndTokenArg struct {
}

type SimpleFSUserEditHistoryArg struct {
}

type SimpleFSFolderEditHistoryArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSListFavoritesArg struct {
}

type SimpleFSGetUserQuotaUsageArg struct {
}

type SimpleFSGetTeamQuotaUsageArg struct {
	TeamName TeamName `codec:"teamName" json:"teamName"`
}

type SimpleFSResetArg struct {
	Path  Path   `codec:"path" json:"path"`
	TlfID string `codec:"tlfID" json:"tlfID"`
}

type SimpleFSFolderSyncConfigAndStatusArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSSetFolderSyncConfigArg struct {
	Path   Path             `codec:"path" json:"path"`
	Config FolderSyncConfig `codec:"config" json:"config"`
}

type SimpleFSSyncConfigAndStatusArg struct {
	IdentifyBehavior *TLFIdentifyBehavior `codec:"identifyBehavior,omitempty" json:"identifyBehavior,omitempty"`
}

type SimpleFSAreWeConnectedToMDServerArg struct {
}

type SimpleFSCheckReachabilityArg struct {
}

type SimpleFSSetDebugLevelArg struct {
	Level string `codec:"level" json:"level"`
}

type SimpleFSSettingsArg struct {
}

type SimpleFSSetNotificationThresholdArg struct {
	Threshold int64 `codec:"threshold" json:"threshold"`
}

type SimpleFSObfuscatePathArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSDeobfuscatePathArg struct {
	Path Path `codec:"path" json:"path"`
}

type SimpleFSGetStatsArg struct {
}

type SimpleFSSubscribePathArg struct {
	IdentifyBehavior          *TLFIdentifyBehavior  `codec:"identifyBehavior,omitempty" json:"identifyBehavior,omitempty"`
	SubscriptionID            string                `codec:"subscriptionID" json:"subscriptionID"`
	KbfsPath                  string                `codec:"kbfsPath" json:"kbfsPath"`
	Topic                     PathSubscriptionTopic `codec:"topic" json:"topic"`
	DeduplicateIntervalSecond int                   `codec:"deduplicateIntervalSecond" json:"deduplicateIntervalSecond"`
}

type SimpleFSSubscribeNonPathArg struct {
	IdentifyBehavior          *TLFIdentifyBehavior `codec:"identifyBehavior,omitempty" json:"identifyBehavior,omitempty"`
	SubscriptionID            string               `codec:"subscriptionID" json:"subscriptionID"`
	Topic                     SubscriptionTopic    `codec:"topic" json:"topic"`
	DeduplicateIntervalSecond int                  `codec:"deduplicateIntervalSecond" json:"deduplicateIntervalSecond"`
}

type SimpleFSUnsubscribeArg struct {
	IdentifyBehavior *TLFIdentifyBehavior `codec:"identifyBehavior,omitempty" json:"identifyBehavior,omitempty"`
	SubscriptionID   string               `codec:"subscriptionID" json:"subscriptionID"`
}

type SimpleFSInterface interface {
	// Begin list of items in directory at path.
	// Retrieve results with readList().
	// Can be a single file to get flags/status.
	// If `refreshSubscription` is true and the path is a KBFS path, simpleFS
	// will begin sending `FSPathUpdated` notifications for the for the
	// corresponding TLF, until another call refreshes the subscription on a
	// different TLF.
	SimpleFSList(context.Context, SimpleFSListArg) error
	// Begin recursive list of items in directory at path.
	// If `refreshSubscription` is true and the path is a KBFS path, simpleFS
	// will begin sending `FSPathUpdated` notifications for the for the
	// corresponding TLF, until another call refreshes the subscription on a
	// different TLF.
	SimpleFSListRecursive(context.Context, SimpleFSListRecursiveArg) error
	// Begin recursive list of items in directory at path up to a given depth
	SimpleFSListRecursiveToDepth(context.Context, SimpleFSListRecursiveToDepthArg) error
	// Get list of Paths in progress. Can indicate status of pending
	// to get more entries.
	SimpleFSReadList(context.Context, OpID) (SimpleFSListResult, error)
	// Begin copy of file or directory
	SimpleFSCopy(context.Context, SimpleFSCopyArg) error
	// Make a symlink of file or directory
	SimpleFSSymlink(context.Context, SimpleFSSymlinkArg) error
	// Begin recursive copy of directory
	SimpleFSCopyRecursive(context.Context, SimpleFSCopyRecursiveArg) error
	// Begin move of file or directory, from/to KBFS only
	SimpleFSMove(context.Context, SimpleFSMoveArg) error
	// Rename file or directory, KBFS side only
	SimpleFSRename(context.Context, SimpleFSRenameArg) error
	// Create/open a file and leave it open
	// or create a directory
	// Files must be closed afterwards.
	SimpleFSOpen(context.Context, SimpleFSOpenArg) error
	// Set/clear file bits - only executable for now
	SimpleFSSetStat(context.Context, SimpleFSSetStatArg) error
	// Read (possibly partial) contents of open file,
	// up to the amount specified by size.
	// Repeat until zero bytes are returned or error.
	// If size is zero, read an arbitrary amount.
	SimpleFSRead(context.Context, SimpleFSReadArg) (FileContent, error)
	// Append content to opened file.
	// May be repeated until OpID is closed.
	SimpleFSWrite(context.Context, SimpleFSWriteArg) error
	// Remove file or directory from filesystem
	SimpleFSRemove(context.Context, SimpleFSRemoveArg) error
	// Get info about file
	SimpleFSStat(context.Context, SimpleFSStatArg) (Dirent, error)
	// Get revision info for a directory entry
	SimpleFSGetRevisions(context.Context, SimpleFSGetRevisionsArg) error
	// Get list of revisions in progress. Can indicate status of pending
	// to get more revisions.
	SimpleFSReadRevisions(context.Context, OpID) (GetRevisionsResult, error)
	// Convenience helper for generating new random value
	SimpleFSMakeOpid(context.Context) (OpID, error)
	// Close OpID, cancels any pending operation.
	// Must be called after list/copy/remove
	SimpleFSClose(context.Context, OpID) error
	// Cancels a running operation, like copy.
	SimpleFSCancel(context.Context, OpID) error
	// Check progress of pending operation
	SimpleFSCheck(context.Context, OpID) (OpProgress, error)
	// Get all the outstanding operations
	SimpleFSGetOps(context.Context) ([]OpDescription, error)
	// Blocking wait for the pending operation to finish
	SimpleFSWait(context.Context, OpID) error
	// Instructs KBFS to dump debugging info into its logs.
	SimpleFSDumpDebuggingInfo(context.Context) error
	SimpleFSClearConflictState(context.Context, Path) error
	SimpleFSFinishResolvingConflict(context.Context, Path) error
	// Force a TLF into a stuck conflict state (for testing).
	SimpleFSForceStuckConflict(context.Context, Path) error
	// Get sync status.
	SimpleFSSyncStatus(context.Context, ListFilter) (FSSyncStatus, error)
	// This RPC generates a random token to be used by a client that needs to
	// access KBFS content through HTTP. It's fine to call this RPC more than once,
	// but it's probably best to call this once and keep using it. Clients should
	// be using HTTP GET methods, with requests in the form of:
	//
	// http://<address>/<kbfs_path>?token=<token>
	//
	// If the provided token is invalid, 403 Forbidden is returned, with a HTML
	// document that has title of "KBFS HTTP Token Invalid". When receiving such
	// response, client should call this RPC (again) to get a new token.
	SimpleFSGetHTTPAddressAndToken(context.Context) (SimpleFSGetHTTPAddressAndTokenResponse, error)
	// simpleFSUserEditHistory returns edit histories of TLFs that the logged-in
	// user can access.  Each returned history is corresponds to a unique
	// writer-TLF pair.  They are in descending order by the modification time
	// (as recorded by the server) of the most recent edit in each history.
	SimpleFSUserEditHistory(context.Context) ([]FSFolderEditHistory, error)
	// simpleFSFolderEditHistory returns the edit history for the TLF
	// described by `path`, for the most recent writers of that TLF.
	// The writers are in descending order by the modification time (as
	// recorded by the server) of their most recent edit.
	SimpleFSFolderEditHistory(context.Context, Path) (FSFolderEditHistory, error)
	// simpleFSListFavorites gets the current favorites, ignored folders, and new
	// folders from the KBFS cache.
	SimpleFSListFavorites(context.Context) (FavoritesResult, error)
	// simpleFSGetUserQuotaUsage returns the quota usage for the logged-in
	// user.  Any usage includes local journal usage as well.
	SimpleFSGetUserQuotaUsage(context.Context) (SimpleFSQuotaUsage, error)
	// simpleFSGetTeamQuotaUsage returns the quota usage for the given team, if
	// the logged-in user has access to that team.  Any usage includes
	// local journal usage as well.
	SimpleFSGetTeamQuotaUsage(context.Context, TeamName) (SimpleFSQuotaUsage, error)
	// simpleFSReset completely resets the KBFS folder referenced in `path`.
	// It should only be called after explicit user confirmation.
	SimpleFSReset(context.Context, SimpleFSResetArg) error
	SimpleFSFolderSyncConfigAndStatus(context.Context, Path) (FolderSyncConfigAndStatus, error)
	SimpleFSSetFolderSyncConfig(context.Context, SimpleFSSetFolderSyncConfigArg) error
	SimpleFSSyncConfigAndStatus(context.Context, *TLFIdentifyBehavior) (SyncConfigAndStatusRes, error)
	SimpleFSAreWeConnectedToMDServer(context.Context) (bool, error)
	SimpleFSCheckReachability(context.Context) error
	SimpleFSSetDebugLevel(context.Context, string) error
	SimpleFSSettings(context.Context) (FSSettings, error)
	SimpleFSSetNotificationThreshold(context.Context, int64) error
	SimpleFSObfuscatePath(context.Context, Path) (string, error)
	SimpleFSDeobfuscatePath(context.Context, Path) ([]string, error)
	SimpleFSGetStats(context.Context) (SimpleFSStats, error)
	SimpleFSSubscribePath(context.Context, SimpleFSSubscribePathArg) error
	SimpleFSSubscribeNonPath(context.Context, SimpleFSSubscribeNonPathArg) error
	SimpleFSUnsubscribe(context.Context, SimpleFSUnsubscribeArg) error
}

func SimpleFSProtocol(i SimpleFSInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.SimpleFS",
		Methods: map[string]rpc.ServeHandlerDescription{
			"simpleFSList": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSListArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSListArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSListArg)(nil), args)
						return
					}
					err = i.SimpleFSList(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSListRecursive": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSListRecursiveArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSListRecursiveArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSListRecursiveArg)(nil), args)
						return
					}
					err = i.SimpleFSListRecursive(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSListRecursiveToDepth": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSListRecursiveToDepthArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSListRecursiveToDepthArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSListRecursiveToDepthArg)(nil), args)
						return
					}
					err = i.SimpleFSListRecursiveToDepth(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSReadList": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSReadListArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSReadListArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSReadListArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSReadList(ctx, typedArgs[0].OpID)
					return
				},
			},
			"simpleFSCopy": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSCopyArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSCopyArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSCopyArg)(nil), args)
						return
					}
					err = i.SimpleFSCopy(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSSymlink": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSymlinkArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSymlinkArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSymlinkArg)(nil), args)
						return
					}
					err = i.SimpleFSSymlink(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSCopyRecursive": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSCopyRecursiveArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSCopyRecursiveArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSCopyRecursiveArg)(nil), args)
						return
					}
					err = i.SimpleFSCopyRecursive(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSMove": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSMoveArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSMoveArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSMoveArg)(nil), args)
						return
					}
					err = i.SimpleFSMove(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSRename": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSRenameArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSRenameArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSRenameArg)(nil), args)
						return
					}
					err = i.SimpleFSRename(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSOpen": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSOpenArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSOpenArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSOpenArg)(nil), args)
						return
					}
					err = i.SimpleFSOpen(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSSetStat": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSetStatArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSetStatArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSetStatArg)(nil), args)
						return
					}
					err = i.SimpleFSSetStat(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSRead": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSReadArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSReadArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSReadArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSRead(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSWrite": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSWriteArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSWriteArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSWriteArg)(nil), args)
						return
					}
					err = i.SimpleFSWrite(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSRemove": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSRemoveArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSRemoveArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSRemoveArg)(nil), args)
						return
					}
					err = i.SimpleFSRemove(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSStat": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSStatArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSStatArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSStatArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSStat(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSGetRevisions": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSGetRevisionsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSGetRevisionsArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSGetRevisionsArg)(nil), args)
						return
					}
					err = i.SimpleFSGetRevisions(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSReadRevisions": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSReadRevisionsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSReadRevisionsArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSReadRevisionsArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSReadRevisions(ctx, typedArgs[0].OpID)
					return
				},
			},
			"simpleFSMakeOpid": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSMakeOpidArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSMakeOpid(ctx)
					return
				},
			},
			"simpleFSClose": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSCloseArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSCloseArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSCloseArg)(nil), args)
						return
					}
					err = i.SimpleFSClose(ctx, typedArgs[0].OpID)
					return
				},
			},
			"simpleFSCancel": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSCancelArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSCancelArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSCancelArg)(nil), args)
						return
					}
					err = i.SimpleFSCancel(ctx, typedArgs[0].OpID)
					return
				},
			},
			"simpleFSCheck": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSCheckArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSCheckArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSCheckArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSCheck(ctx, typedArgs[0].OpID)
					return
				},
			},
			"simpleFSGetOps": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSGetOpsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSGetOps(ctx)
					return
				},
			},
			"simpleFSWait": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSWaitArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSWaitArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSWaitArg)(nil), args)
						return
					}
					err = i.SimpleFSWait(ctx, typedArgs[0].OpID)
					return
				},
			},
			"simpleFSDumpDebuggingInfo": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSDumpDebuggingInfoArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.SimpleFSDumpDebuggingInfo(ctx)
					return
				},
			},
			"simpleFSClearConflictState": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSClearConflictStateArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSClearConflictStateArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSClearConflictStateArg)(nil), args)
						return
					}
					err = i.SimpleFSClearConflictState(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSFinishResolvingConflict": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSFinishResolvingConflictArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSFinishResolvingConflictArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSFinishResolvingConflictArg)(nil), args)
						return
					}
					err = i.SimpleFSFinishResolvingConflict(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSForceStuckConflict": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSForceStuckConflictArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSForceStuckConflictArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSForceStuckConflictArg)(nil), args)
						return
					}
					err = i.SimpleFSForceStuckConflict(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSSyncStatus": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSyncStatusArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSyncStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSyncStatusArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSSyncStatus(ctx, typedArgs[0].Filter)
					return
				},
			},
			"SimpleFSGetHTTPAddressAndToken": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSGetHTTPAddressAndTokenArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSGetHTTPAddressAndToken(ctx)
					return
				},
			},
			"simpleFSUserEditHistory": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSUserEditHistoryArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSUserEditHistory(ctx)
					return
				},
			},
			"simpleFSFolderEditHistory": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSFolderEditHistoryArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSFolderEditHistoryArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSFolderEditHistoryArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSFolderEditHistory(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSListFavorites": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSListFavoritesArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSListFavorites(ctx)
					return
				},
			},
			"simpleFSGetUserQuotaUsage": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSGetUserQuotaUsageArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSGetUserQuotaUsage(ctx)
					return
				},
			},
			"simpleFSGetTeamQuotaUsage": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSGetTeamQuotaUsageArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSGetTeamQuotaUsageArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSGetTeamQuotaUsageArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSGetTeamQuotaUsage(ctx, typedArgs[0].TeamName)
					return
				},
			},
			"simpleFSReset": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSResetArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSResetArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSResetArg)(nil), args)
						return
					}
					err = i.SimpleFSReset(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSFolderSyncConfigAndStatus": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSFolderSyncConfigAndStatusArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSFolderSyncConfigAndStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSFolderSyncConfigAndStatusArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSFolderSyncConfigAndStatus(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSSetFolderSyncConfig": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSetFolderSyncConfigArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSetFolderSyncConfigArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSetFolderSyncConfigArg)(nil), args)
						return
					}
					err = i.SimpleFSSetFolderSyncConfig(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSSyncConfigAndStatus": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSyncConfigAndStatusArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSyncConfigAndStatusArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSyncConfigAndStatusArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSSyncConfigAndStatus(ctx, typedArgs[0].IdentifyBehavior)
					return
				},
			},
			"simpleFSAreWeConnectedToMDServer": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSAreWeConnectedToMDServerArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSAreWeConnectedToMDServer(ctx)
					return
				},
			},
			"simpleFSCheckReachability": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSCheckReachabilityArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.SimpleFSCheckReachability(ctx)
					return
				},
			},
			"simpleFSSetDebugLevel": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSetDebugLevelArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSetDebugLevelArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSetDebugLevelArg)(nil), args)
						return
					}
					err = i.SimpleFSSetDebugLevel(ctx, typedArgs[0].Level)
					return
				},
			},
			"simpleFSSettings": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSettingsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSSettings(ctx)
					return
				},
			},
			"simpleFSSetNotificationThreshold": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSetNotificationThresholdArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSetNotificationThresholdArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSetNotificationThresholdArg)(nil), args)
						return
					}
					err = i.SimpleFSSetNotificationThreshold(ctx, typedArgs[0].Threshold)
					return
				},
			},
			"simpleFSObfuscatePath": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSObfuscatePathArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSObfuscatePathArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSObfuscatePathArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSObfuscatePath(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSDeobfuscatePath": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSDeobfuscatePathArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSDeobfuscatePathArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSDeobfuscatePathArg)(nil), args)
						return
					}
					ret, err = i.SimpleFSDeobfuscatePath(ctx, typedArgs[0].Path)
					return
				},
			},
			"simpleFSGetStats": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSGetStatsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.SimpleFSGetStats(ctx)
					return
				},
			},
			"simpleFSSubscribePath": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSubscribePathArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSubscribePathArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSubscribePathArg)(nil), args)
						return
					}
					err = i.SimpleFSSubscribePath(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSSubscribeNonPath": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSSubscribeNonPathArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSSubscribeNonPathArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSSubscribeNonPathArg)(nil), args)
						return
					}
					err = i.SimpleFSSubscribeNonPath(ctx, typedArgs[0])
					return
				},
			},
			"simpleFSUnsubscribe": {
				MakeArg: func() interface{} {
					var ret [1]SimpleFSUnsubscribeArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SimpleFSUnsubscribeArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SimpleFSUnsubscribeArg)(nil), args)
						return
					}
					err = i.SimpleFSUnsubscribe(ctx, typedArgs[0])
					return
				},
			},
		},
	}
}

type SimpleFSClient struct {
	Cli rpc.GenericClient
}

// Begin list of items in directory at path.
// Retrieve results with readList().
// Can be a single file to get flags/status.
// If `refreshSubscription` is true and the path is a KBFS path, simpleFS
// will begin sending `FSPathUpdated` notifications for the for the
// corresponding TLF, until another call refreshes the subscription on a
// different TLF.
func (c SimpleFSClient) SimpleFSList(ctx context.Context, __arg SimpleFSListArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSList", []interface{}{__arg}, nil)
	return
}

// Begin recursive list of items in directory at path.
// If `refreshSubscription` is true and the path is a KBFS path, simpleFS
// will begin sending `FSPathUpdated` notifications for the for the
// corresponding TLF, until another call refreshes the subscription on a
// different TLF.
func (c SimpleFSClient) SimpleFSListRecursive(ctx context.Context, __arg SimpleFSListRecursiveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSListRecursive", []interface{}{__arg}, nil)
	return
}

// Begin recursive list of items in directory at path up to a given depth
func (c SimpleFSClient) SimpleFSListRecursiveToDepth(ctx context.Context, __arg SimpleFSListRecursiveToDepthArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSListRecursiveToDepth", []interface{}{__arg}, nil)
	return
}

// Get list of Paths in progress. Can indicate status of pending
// to get more entries.
func (c SimpleFSClient) SimpleFSReadList(ctx context.Context, opID OpID) (res SimpleFSListResult, err error) {
	__arg := SimpleFSReadListArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSReadList", []interface{}{__arg}, &res)
	return
}

// Begin copy of file or directory
func (c SimpleFSClient) SimpleFSCopy(ctx context.Context, __arg SimpleFSCopyArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCopy", []interface{}{__arg}, nil)
	return
}

// Make a symlink of file or directory
func (c SimpleFSClient) SimpleFSSymlink(ctx context.Context, __arg SimpleFSSymlinkArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSymlink", []interface{}{__arg}, nil)
	return
}

// Begin recursive copy of directory
func (c SimpleFSClient) SimpleFSCopyRecursive(ctx context.Context, __arg SimpleFSCopyRecursiveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCopyRecursive", []interface{}{__arg}, nil)
	return
}

// Begin move of file or directory, from/to KBFS only
func (c SimpleFSClient) SimpleFSMove(ctx context.Context, __arg SimpleFSMoveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSMove", []interface{}{__arg}, nil)
	return
}

// Rename file or directory, KBFS side only
func (c SimpleFSClient) SimpleFSRename(ctx context.Context, __arg SimpleFSRenameArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSRename", []interface{}{__arg}, nil)
	return
}

// Create/open a file and leave it open
// or create a directory
// Files must be closed afterwards.
func (c SimpleFSClient) SimpleFSOpen(ctx context.Context, __arg SimpleFSOpenArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSOpen", []interface{}{__arg}, nil)
	return
}

// Set/clear file bits - only executable for now
func (c SimpleFSClient) SimpleFSSetStat(ctx context.Context, __arg SimpleFSSetStatArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSetStat", []interface{}{__arg}, nil)
	return
}

// Read (possibly partial) contents of open file,
// up to the amount specified by size.
// Repeat until zero bytes are returned or error.
// If size is zero, read an arbitrary amount.
func (c SimpleFSClient) SimpleFSRead(ctx context.Context, __arg SimpleFSReadArg) (res FileContent, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSRead", []interface{}{__arg}, &res)
	return
}

// Append content to opened file.
// May be repeated until OpID is closed.
func (c SimpleFSClient) SimpleFSWrite(ctx context.Context, __arg SimpleFSWriteArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSWrite", []interface{}{__arg}, nil)
	return
}

// Remove file or directory from filesystem
func (c SimpleFSClient) SimpleFSRemove(ctx context.Context, __arg SimpleFSRemoveArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSRemove", []interface{}{__arg}, nil)
	return
}

// Get info about file
func (c SimpleFSClient) SimpleFSStat(ctx context.Context, __arg SimpleFSStatArg) (res Dirent, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSStat", []interface{}{__arg}, &res)
	return
}

// Get revision info for a directory entry
func (c SimpleFSClient) SimpleFSGetRevisions(ctx context.Context, __arg SimpleFSGetRevisionsArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSGetRevisions", []interface{}{__arg}, nil)
	return
}

// Get list of revisions in progress. Can indicate status of pending
// to get more revisions.
func (c SimpleFSClient) SimpleFSReadRevisions(ctx context.Context, opID OpID) (res GetRevisionsResult, err error) {
	__arg := SimpleFSReadRevisionsArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSReadRevisions", []interface{}{__arg}, &res)
	return
}

// Convenience helper for generating new random value
func (c SimpleFSClient) SimpleFSMakeOpid(ctx context.Context) (res OpID, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSMakeOpid", []interface{}{SimpleFSMakeOpidArg{}}, &res)
	return
}

// Close OpID, cancels any pending operation.
// Must be called after list/copy/remove
func (c SimpleFSClient) SimpleFSClose(ctx context.Context, opID OpID) (err error) {
	__arg := SimpleFSCloseArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSClose", []interface{}{__arg}, nil)
	return
}

// Cancels a running operation, like copy.
func (c SimpleFSClient) SimpleFSCancel(ctx context.Context, opID OpID) (err error) {
	__arg := SimpleFSCancelArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCancel", []interface{}{__arg}, nil)
	return
}

// Check progress of pending operation
func (c SimpleFSClient) SimpleFSCheck(ctx context.Context, opID OpID) (res OpProgress, err error) {
	__arg := SimpleFSCheckArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCheck", []interface{}{__arg}, &res)
	return
}

// Get all the outstanding operations
func (c SimpleFSClient) SimpleFSGetOps(ctx context.Context) (res []OpDescription, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSGetOps", []interface{}{SimpleFSGetOpsArg{}}, &res)
	return
}

// Blocking wait for the pending operation to finish
func (c SimpleFSClient) SimpleFSWait(ctx context.Context, opID OpID) (err error) {
	__arg := SimpleFSWaitArg{OpID: opID}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSWait", []interface{}{__arg}, nil)
	return
}

// Instructs KBFS to dump debugging info into its logs.
func (c SimpleFSClient) SimpleFSDumpDebuggingInfo(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSDumpDebuggingInfo", []interface{}{SimpleFSDumpDebuggingInfoArg{}}, nil)
	return
}

func (c SimpleFSClient) SimpleFSClearConflictState(ctx context.Context, path Path) (err error) {
	__arg := SimpleFSClearConflictStateArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSClearConflictState", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSFinishResolvingConflict(ctx context.Context, path Path) (err error) {
	__arg := SimpleFSFinishResolvingConflictArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSFinishResolvingConflict", []interface{}{__arg}, nil)
	return
}

// Force a TLF into a stuck conflict state (for testing).
func (c SimpleFSClient) SimpleFSForceStuckConflict(ctx context.Context, path Path) (err error) {
	__arg := SimpleFSForceStuckConflictArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSForceStuckConflict", []interface{}{__arg}, nil)
	return
}

// Get sync status.
func (c SimpleFSClient) SimpleFSSyncStatus(ctx context.Context, filter ListFilter) (res FSSyncStatus, err error) {
	__arg := SimpleFSSyncStatusArg{Filter: filter}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSyncStatus", []interface{}{__arg}, &res)
	return
}

// This RPC generates a random token to be used by a client that needs to
// access KBFS content through HTTP. It's fine to call this RPC more than once,
// but it's probably best to call this once and keep using it. Clients should
// be using HTTP GET methods, with requests in the form of:
//
// http://<address>/<kbfs_path>?token=<token>
//
// If the provided token is invalid, 403 Forbidden is returned, with a HTML
// document that has title of "KBFS HTTP Token Invalid". When receiving such
// response, client should call this RPC (again) to get a new token.
func (c SimpleFSClient) SimpleFSGetHTTPAddressAndToken(ctx context.Context) (res SimpleFSGetHTTPAddressAndTokenResponse, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.SimpleFSGetHTTPAddressAndToken", []interface{}{SimpleFSGetHTTPAddressAndTokenArg{}}, &res)
	return
}

// simpleFSUserEditHistory returns edit histories of TLFs that the logged-in
// user can access.  Each returned history is corresponds to a unique
// writer-TLF pair.  They are in descending order by the modification time
// (as recorded by the server) of the most recent edit in each history.
func (c SimpleFSClient) SimpleFSUserEditHistory(ctx context.Context) (res []FSFolderEditHistory, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSUserEditHistory", []interface{}{SimpleFSUserEditHistoryArg{}}, &res)
	return
}

// simpleFSFolderEditHistory returns the edit history for the TLF
// described by `path`, for the most recent writers of that TLF.
// The writers are in descending order by the modification time (as
// recorded by the server) of their most recent edit.
func (c SimpleFSClient) SimpleFSFolderEditHistory(ctx context.Context, path Path) (res FSFolderEditHistory, err error) {
	__arg := SimpleFSFolderEditHistoryArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSFolderEditHistory", []interface{}{__arg}, &res)
	return
}

// simpleFSListFavorites gets the current favorites, ignored folders, and new
// folders from the KBFS cache.
func (c SimpleFSClient) SimpleFSListFavorites(ctx context.Context) (res FavoritesResult, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSListFavorites", []interface{}{SimpleFSListFavoritesArg{}}, &res)
	return
}

// simpleFSGetUserQuotaUsage returns the quota usage for the logged-in
// user.  Any usage includes local journal usage as well.
func (c SimpleFSClient) SimpleFSGetUserQuotaUsage(ctx context.Context) (res SimpleFSQuotaUsage, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSGetUserQuotaUsage", []interface{}{SimpleFSGetUserQuotaUsageArg{}}, &res)
	return
}

// simpleFSGetTeamQuotaUsage returns the quota usage for the given team, if
// the logged-in user has access to that team.  Any usage includes
// local journal usage as well.
func (c SimpleFSClient) SimpleFSGetTeamQuotaUsage(ctx context.Context, teamName TeamName) (res SimpleFSQuotaUsage, err error) {
	__arg := SimpleFSGetTeamQuotaUsageArg{TeamName: teamName}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSGetTeamQuotaUsage", []interface{}{__arg}, &res)
	return
}

// simpleFSReset completely resets the KBFS folder referenced in `path`.
// It should only be called after explicit user confirmation.
func (c SimpleFSClient) SimpleFSReset(ctx context.Context, __arg SimpleFSResetArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSReset", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSFolderSyncConfigAndStatus(ctx context.Context, path Path) (res FolderSyncConfigAndStatus, err error) {
	__arg := SimpleFSFolderSyncConfigAndStatusArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSFolderSyncConfigAndStatus", []interface{}{__arg}, &res)
	return
}

func (c SimpleFSClient) SimpleFSSetFolderSyncConfig(ctx context.Context, __arg SimpleFSSetFolderSyncConfigArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSetFolderSyncConfig", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSSyncConfigAndStatus(ctx context.Context, identifyBehavior *TLFIdentifyBehavior) (res SyncConfigAndStatusRes, err error) {
	__arg := SimpleFSSyncConfigAndStatusArg{IdentifyBehavior: identifyBehavior}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSyncConfigAndStatus", []interface{}{__arg}, &res)
	return
}

func (c SimpleFSClient) SimpleFSAreWeConnectedToMDServer(ctx context.Context) (res bool, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSAreWeConnectedToMDServer", []interface{}{SimpleFSAreWeConnectedToMDServerArg{}}, &res)
	return
}

func (c SimpleFSClient) SimpleFSCheckReachability(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSCheckReachability", []interface{}{SimpleFSCheckReachabilityArg{}}, nil)
	return
}

func (c SimpleFSClient) SimpleFSSetDebugLevel(ctx context.Context, level string) (err error) {
	__arg := SimpleFSSetDebugLevelArg{Level: level}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSetDebugLevel", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSSettings(ctx context.Context) (res FSSettings, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSettings", []interface{}{SimpleFSSettingsArg{}}, &res)
	return
}

func (c SimpleFSClient) SimpleFSSetNotificationThreshold(ctx context.Context, threshold int64) (err error) {
	__arg := SimpleFSSetNotificationThresholdArg{Threshold: threshold}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSetNotificationThreshold", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSObfuscatePath(ctx context.Context, path Path) (res string, err error) {
	__arg := SimpleFSObfuscatePathArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSObfuscatePath", []interface{}{__arg}, &res)
	return
}

func (c SimpleFSClient) SimpleFSDeobfuscatePath(ctx context.Context, path Path) (res []string, err error) {
	__arg := SimpleFSDeobfuscatePathArg{Path: path}
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSDeobfuscatePath", []interface{}{__arg}, &res)
	return
}

func (c SimpleFSClient) SimpleFSGetStats(ctx context.Context) (res SimpleFSStats, err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSGetStats", []interface{}{SimpleFSGetStatsArg{}}, &res)
	return
}

func (c SimpleFSClient) SimpleFSSubscribePath(ctx context.Context, __arg SimpleFSSubscribePathArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSubscribePath", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSSubscribeNonPath(ctx context.Context, __arg SimpleFSSubscribeNonPathArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSSubscribeNonPath", []interface{}{__arg}, nil)
	return
}

func (c SimpleFSClient) SimpleFSUnsubscribe(ctx context.Context, __arg SimpleFSUnsubscribeArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.SimpleFS.simpleFSUnsubscribe", []interface{}{__arg}, nil)
	return
}
