// Поганий приклад
package yamlconfig

func ParseYAMLConfig(input string) (*Config, error)

// Гарний приклад
package yamlconfig

func Parse(input string) (*Config, error)



// Поганий приклад
func (c *Config) GetJobName(key string) (value string, ok bool)

// Гарний приклад
func (c *Config) JobName(key string) (value string, ok bool)



// Поганий приклад
const MAX_PACKET_SIZE = 512
const kMaxBufferSize = 1024

// Гарний приклад
const MaxPacketSize = 512
const MaxBufferSize = 1024



// Поганий приклад
func GetUrlInfo(id int) {}

// Гарний приклад
func GetURLInfo(id int) {}



// Поганий приклад
if strings.Contains(err.Error(), "duplicate") { ... }

// Гарний приклад
if errors.Is(err, ErrDuplicate) { ... }



// Поганий приклад
if user, err = db.UserByID(userID); err != nil {
    // ...
}

// Гарний приклад
u, err := db.UserByID(userID)
if err != nil {
    return fmt.Errorf("invalid origin user: %s", err)
}
user = u


// Поганий приклад
// This is a comment paragraph. The length of individual lines doesn't matter in Godoc;
// but the choice of wrapping causes jagged lines on narrow screens or in code review.

// Гарний приклад
// This is a comment paragraph.
// Wrapping comments improves readability
// on narrow screens and in code reviews.



// Поганий приклад
func Bad() *os.PathError {
    // ...
}

// Гарний приклад
func GoodLookup() (*Result, error) {
    if err != nil {
        return nil, err
    }
    return res, nil
}



// Поганий приклад
m := make(map[string]os.DirEntry)

// Гарний приклад
files, _ := os.ReadDir("./files")
m := make(map[string]os.DirEntry, len(files))



// Поганий приклад
for i := 0; i < b.N; i++ {
    w.Write([]byte("Hello world"))
}

// Гарний приклад
data := []byte("Hello world")
for i := 0; i < b.N; i++ {
    w.Write(data)
}



// Поганий приклад
for _, v := range data {
    if v.F1 == 1 {
        v = process(v)
        if err := v.Call(); err == nil {
            v.Send()
        } else {
            return err
        }
    } else {
        log.Printf("Invalid v: %v", v)
    }
}

// Гарний приклад
for _, v := range data {
    if v.F1 != 1 {
        log.Printf("Invalid v: %v", v)
        continue
    }

    v = process(v)
    if err := v.Call(); err != nil {
        return err
    }
    v.Send()
}



// Поганий приклад
const (
    defaultPort = 8080
    defaultUser = "user"
)

// Гарний приклад
const (
    _defaultPort = 8080
    _defaultUser = "user"
)



// Поганий приклад
func f(list []int) {
    filtered := []int{}
    for _, v := range list {
        if v > 10 {
            filtered = append(filtered, v)
        }
    }
}

// Гарний приклад
func f(list []int) {
    var filtered []int
    for _, v := range list {
        if v > 10 {
            filtered = append(filtered, v)
        }
    }
}



// Поганий приклад
func Open(addr string, cache bool, logger *zap.Logger) (*Connection, error) {
    // ...
}

// Гарний приклад
func Open(addr string, opts ...Option) (*Connection, error) {
    options := options{
        cache:  defaultCache,
        logger: zap.NewNop(),
    }

    for _, o := range opts {
        o.apply(&options)
    }

    // ...
    return &Connection{}, nil
}
